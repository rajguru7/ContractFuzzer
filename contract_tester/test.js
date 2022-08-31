const Web3 = require("web3");
const Promise = require("bluebird");

const accounts = [
    "0x2b71cc952c8e3dfe97a696cf5c5b29f8a07de3d8",
    "0xa31a0f4653f62aca35b6e986743c8f4fc6c8f38f",
    "0x6d62f53305d3c247cd856a8a4eaf65518a7030cf",
    "0x04c8862a82faf3fb90b73768c50dc7f23d7d26bd",
    "0x271eab2a8058d1e3a45941f49d5671bb1cee8ca1"
];
const defaultAccount = accounts[0];
const Account1 = accounts[1];

let Provider;
let web3;
const httpRpcAddr = "http://localhost:8545"
Provider = new Web3.providers.HttpProvider(httpRpcAddr);
web3 = new Web3(Provider);
web3.personal.unlockAccount(defaultAccount, "123456", 200 * 60 * 60);
web3.personal.unlockAccount(Account1, "123456", 200 * 60 * 60);

const sendTransaction = Promise.promisify(web3.eth.sendTransaction);

let transaction = {from: "0x2b71cc952c8e3dfe97a696cf5c5b29f8a07de3d8", to: "0x30a6e91692d091e904cd4d9b7c9895dbdd6711de", value: "10000000000", gas: "800000", data: "0x0c55699c"};
console.log("sending transaction \n");
sendTransaction(transaction).catch(function(err){
    //do nothing but output err msg
    console.log(err);            
});
