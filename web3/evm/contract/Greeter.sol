//SPDX-License-Identifier: Unlicense
pragma solidity ^0.8.0;

// import "hardhat/console.sol";

contract Greeter {
    event GreeterChange( string indexed greeting, uint256 version);
    string private greeting;
    uint256 private version;


    constructor(string memory _greeting) {
        // console.log("Deploying a Greeter with greeting:", _greeting);
        greeting = _greeting;
        version = 1;
    }

    function greet() public view returns (string memory) {
        return greeting;
    }

    function setGreeting(string memory _greeting) public {
        // console.log("Changing greeting from '%s' to '%s'", greeting, _greeting);
        greeting = _greeting;
        version += 1;
        emit GreeterChange(_greeting, version);
    }

    function getOverview() public view returns (string memory, uint256) {
        return (greeting, version);
    }
}
