// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.9;

import "./IClient.sol";
import "./IBCClient.sol";
import "./IBCChannel.sol";
import "./IBCModule.sol";
import "./IBCMsgs.sol";
import "./IBCIdentifier.sol";

contract IBCHandler {

    address owner;
    IBCHost host;
    uint256[] steps;

    /// Event definitions ///
    event SendPacket(Packet.Data packet);
    event RecvPacket(Packet.Data packet);
    event WriteAcknowledgement(string destinationPortId, string destinationChannel, uint64 sequence, bytes acknowledgement);
    event AcknowledgePacket(Packet.Data packet, bytes acknowledgement);

    constructor(IBCHost host_) {
        owner = msg.sender;
        host = host_;
    }

    function getHostAddress() external view returns (address) {
        return address(host);
    }

    function registerClient(string calldata clientType, IClient client) external {
        require(msg.sender == owner);
        return IBCClient.registerClient(host, clientType, client);
    }

    /// Handler interface implementations ///

    function createClient(IBCMsgs.MsgCreateClient calldata msg_) external {
        return IBCClient.createClient(host, msg_);
    }

    function updateClient(IBCMsgs.MsgUpdateClient calldata msg_) external {
        return IBCClient.updateClient(host, msg_);
    }

    function connectionOpenInit(IBCMsgs.MsgConnectionOpenInit memory msg_) public returns (string memory) {
        return IBCConnection.connectionOpenInit(host, msg_);
    }

    function connectionOpenTry(IBCMsgs.MsgConnectionOpenTry memory msg_) public returns (string memory) {
        return IBCConnection.connectionOpenTry(host, msg_);
    }

    function connectionOpenAck(IBCMsgs.MsgConnectionOpenAck memory msg_) public {
        return IBCConnection.connectionOpenAck(host, msg_);
    }

    function connectionOpenConfirm(IBCMsgs.MsgConnectionOpenConfirm memory msg_) public {
        return IBCConnection.connectionOpenConfirm(host, msg_);
    }

    function channelOpenInit(IBCMsgs.MsgChannelOpenInit memory msg_) public returns (string memory) {
        string memory channelId = IBCChannel.channelOpenInit(host, msg_);
        IModuleCallbacks module = lookupModuleByPortId(msg_.portId);
        module.onChanOpenInit(
            msg_.channel.ordering,
            msg_.channel.connection_hops,
            msg_.portId,
            channelId,
            msg_.channel.counterparty,
            msg_.channel.version
        );
        host.claimCapability(IBCIdentifier.channelCapabilityPath(msg_.portId, channelId), address(module));
        return channelId;
    }

    function channelOpenTry(IBCMsgs.MsgChannelOpenTry memory msg_) public returns (string memory) {
        string memory channelId = IBCChannel.channelOpenTry(host, msg_);
        IModuleCallbacks module = lookupModuleByPortId(msg_.portId);
        module.onChanOpenTry(
            msg_.channel.ordering,
            msg_.channel.connection_hops,
            msg_.portId,
            channelId,
            msg_.channel.counterparty,
            msg_.channel.version,
            msg_.counterpartyVersion
        );
        host.claimCapability(IBCIdentifier.channelCapabilityPath(msg_.portId, channelId), address(module));
        return channelId;
    }

    function channelOpenAck(IBCMsgs.MsgChannelOpenAck memory msg_) public {
        IBCChannel.channelOpenAck(host, msg_);
        lookupModuleByPortId(msg_.portId).onChanOpenAck(msg_.portId, msg_.channelId, msg_.counterpartyVersion);
    }

    function channelOpenConfirm(IBCMsgs.MsgChannelOpenConfirm memory msg_) public {
        IBCChannel.channelOpenConfirm(host, msg_);
        lookupModuleByPortId(msg_.portId).onChanOpenConfirm(msg_.portId, msg_.channelId);
    }

    function channelCloseInit(IBCMsgs.MsgChannelCloseInit memory msg_) public {
        IBCChannel.channelCloseInit(host, msg_);
        lookupModuleByPortId(msg_.portId).onChanCloseInit(msg_.portId, msg_.channelId);
    }

    function channelCloseConfirm(IBCMsgs.MsgChannelCloseConfirm memory msg_) public {
        IBCChannel.channelCloseConfirm(host, msg_);
        lookupModuleByPortId(msg_.portId).onChanCloseConfirm(msg_.portId, msg_.channelId);
    }

    function sendPacket(Packet.Data calldata packet) external {
        require(host.authenticateCapability(
            IBCIdentifier.channelCapabilityPath(packet.source_port, packet.source_channel),
            msg.sender
        ));
        IBCChannel.sendPacket(host, packet);
        emit SendPacket(packet);
    }

    function recvPacket(IBCMsgs.MsgPacketRecv calldata msg_) external returns (bytes memory acknowledgement) {
        steps.push(1);
        IModuleCallbacks module = lookupModuleByChannel(msg_.packet.destination_port, msg_.packet.destination_channel);
        steps.push(5);
        acknowledgement = module.onRecvPacket(msg_.packet);
        steps.push(12);
        IBCChannel.recvPacket(host, msg_);
        steps.push(16);
        if (acknowledgement.length > 0) {
            steps.push(17);
            IBCChannel.writeAcknowledgement(host, msg_.packet.destination_port, msg_.packet.destination_channel, msg_.packet.sequence, acknowledgement);
            emit WriteAcknowledgement(msg_.packet.destination_port, msg_.packet.destination_channel, msg_.packet.sequence, acknowledgement);
        }
        emit RecvPacket(msg_.packet);
        return acknowledgement;
    }

    function writeAcknowledgement(string calldata destinationPortId, string calldata destinationChannel, uint64 sequence, bytes calldata acknowledgement) external {
        require(host.authenticateCapability(
            IBCIdentifier.channelCapabilityPath(destinationPortId, destinationChannel),
            msg.sender
        ));
        IBCChannel.writeAcknowledgement(host, destinationPortId, destinationChannel, sequence, acknowledgement);
        emit WriteAcknowledgement(destinationPortId, destinationChannel, sequence, acknowledgement);
    }

    function acknowledgePacket(IBCMsgs.MsgPacketAcknowledgement calldata msg_) external {
        IModuleCallbacks module = lookupModuleByChannel(msg_.packet.source_port, msg_.packet.source_channel);
        module.onAcknowledgementPacket(msg_.packet, msg_.acknowledgement);
        IBCChannel.acknowledgePacket(host, msg_);
        emit AcknowledgePacket(msg_.packet, msg_.acknowledgement);
    }

    function bindPort(string memory portId, address moduleAddress) public {
        onlyOwner();
        host.claimCapability(IBCIdentifier.portCapabilityPath(portId), moduleAddress);
    }

    function setExpectedTimePerBlock(uint64 expectedTimePerBlock_) external {
        // TODO: consider better authn/authz for this operation
        onlyOwner();
        host.setExpectedTimePerBlock(expectedTimePerBlock_);
    }

    function lookupModuleByPortId(string memory portId) internal view returns (IModuleCallbacks) {
        (address module, bool found) = host.getModuleOwner(IBCIdentifier.portCapabilityPath(portId));
        require(found, "not found module owner by port id");
        return IModuleCallbacks(module);
    }

    function lookupModuleByChannel(string memory portId, string memory channelId) public returns (IModuleCallbacks) {
        steps.push(2);
        (address module, bool found) = host.getModuleOwner(IBCIdentifier.channelCapabilityPath(portId, channelId));
        steps.push(3);
        require(found, "not found module owner by channel");
        steps.push(4);
        return IModuleCallbacks(module);
    }

    function onlyOwner() internal view {
        require(msg.sender == owner, "only handler owner");
    }

    function version() public pure returns (string memory) {
        return "v0.0.18";
    }

    function height(uint256 number) public view returns (uint256) {
        require(number != 0, "invalid number");
        return block.number;
    }

    function stepsLength() public view returns (uint256) {
        return steps.length;
    }

    function pushStep(uint256 s) public {
        return steps.push(s);
    }

    function getStep(uint256 index) public view returns (uint256) {
        return steps[index];
    }

    function getSteps() public view returns (uint256[] memory) {
        return steps;
    }
}
