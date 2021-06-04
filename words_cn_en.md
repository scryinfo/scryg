#开发常用命名中英对照
##常见名字
en|cn|note
----|----|----
cashbox|cashbox|
block|区块|
block chain|区块链|
block number|区块号|也有一种是说区块高度
ethereum|以太坊|简称eth
eth|eth|ethereum的token的symbol
bitcoin|比特币|简称btc
btc|btc|bitcoin的token的symbol
bsv|bsv|bitcoin sv(Satoshi Vision)，不要使用中文
bch|bch|bitcoin cash，比特现金
evm|evm|Ethereum Virtual Machine,不建议翻译中文
jvm|java虚拟机|
wallet|钱包|
consensus|共识|
mnemonic|助记词|
transaction|交易|
tx|交易|transaction的缩写
token|token|通证，直接使用英文不翻译
eip|eip|Ethereum Request For Comment/以太坊意见征求稿，erc的范围更大
erc|erc|Ethereum Improvement Proposals/以太坊改进建议
rfc|rfc|请求意见稿（英语：Request for Comments，缩写：RFC）
erc20|erc20|
erc223|erc223|
erc721|erc721|
erc1155|erc1155|
symbol|symbol|token的symbol，如ddd等
decimals|精度|就是小数后的位数，注意有s
balance|余额|
transfer|转帐|
approve| |允许，在token中特指，允许指定帐号地址从自己转帐的数量
value| |如果使用在ethereum中，特指eth的数量
account| |在ethereum中是地址
contract|智能合约|
miner|矿工|
mining machine|矿机|
Segwit|隔离见证|Segregated Witness简称Segwit
Taproot|Taproot|In January 2018, Bitcoin developer Gregory Maxwell proposed a new BTC protocol update, called ‘Taproot’. According to his paper, the new technology will increase transaction privacy, make it more efficient, and eliminate some of SegWit’s deficiencies
slice|动态数组|在go与rust中的概念
array|数组|在go与rust中都是确定大小的，在一些语言中array也是动态的
vector|vector|也是一种动态数组的容器，在c++，rust中使用，不要翻译成中文。
std|标准"库"|它是standard的简称，但翻译成中文叫标准库，虽然单词中没有lib
sdk|sdk |software development kit，软件开发工具包，不建议翻译为中文
jdk|jdk|java development kit，java开发工具包，不建议翻译为中文

##方法名，一般为动词
en| cn |note
----|----|----
create|创建|如果new为语言的关键字，使用create命名，不能同时间使用new与create
new|构造|如果new为语言的关键字，使用create命名，不能同时间使用new与create
add|添加|add主要用于增加/加入一个
remove|删除|
update|更新|有修改之意，不要使用在升级上
edit|修改|update与edit都有修改之意，同一个对象不要同时出现方法update与edit
upgrade|升级|
get|取得|单条记录
list|列表|多条记录
page|分页|
count|计数|也可以说统计

##正反词
en|cn|node
----|---|----
add/remove|添加/删除|不要与insert/delete混用
insert/delete|插入/删除|不要与add/remove混用
open/close|打开/关闭
begin/end| | 
start/stop| | 
show/hide| |
create/destroy| |
source/target| |
first/last| |
min/max| |
get/set| |
up/down| |
old/new| |
next/previous| |

##密码加密或签名
en|cn|note
----|----|----
DES|DES|数据加密标准（英语：Data Encryption Standard，缩写为 DES），改用AES
AES|AES|高级加密标准（英语：Advanced Encryption Standard，缩写：AES），是用于替代DES的
ECC|椭圆曲线密码学|Elliptic Curve Cryptography
RSA|RSA加密算法|是三个人名的简称（Rivest-Shamir-Adleman），是一种非对称加密算法
DSA|DSA|数字签名算法（Digital Signature Algorithm，DSA）
ECDSA|ECDSA|椭圆曲线数字签名算法（英语：Elliptic Curve Digital Signature Algorithm，缩写：ECDSA），是基于Weierstrass 曲线
EdDSA|EdDSA|Edwards-curve Digital Signature Algorithm (EdDSA)，是基于Edwards曲线，第二个字母d是小写。
signature|签名|名词，表示签名后的数据
sign|签名|动词，表示签名的动着
verify| |验证签名，当使用在签名时
secp256k1|secp256k1|NIST系列曲线之一，被用来做签名算法，属于ECDSA，这条曲线也可以用于ECDH
ed25519|ed25519|25519椭圆曲线的一种，爱德华曲线（Edwards Curve），被用来做签名算法，属于EdDSA
curve25519/x25519|curve25519|25519椭圆曲线的一种，蒙哥马利曲线（Montgomery Curve），被用来做ECDH的
ECDH|ECDH|椭圆曲线迪菲-赫尔曼密钥交换（英语：Elliptic Curve Diffie–Hellman key exchange，缩写为ECDH）
DH|DH|迪菲-赫尔曼密钥交换（英语：Diffie–Hellman key exchange，缩写为D-H）
Forward Secrecy| |前向保密（英语：Forward Secrecy，FS）
Perfect Forward Secrecy| |完全前向保密（英语：Perfect Forward Secrecy，PFS）
DHE|DHE|迪菲-赫尔曼密钥交换（DHE，DH Ephemeral）的前向安全通讯
ECDHE|ECDHE|基于椭圆曲线迪菲-赫尔曼密钥交换（ECDHE，ECDH Ephemeral）的前向安全通讯
Schnorr|Schnorr|
MuSig|MuSig|MuSig：由Blockstream提出的Schnorr签名方案
BLS|BLS|Boneh–Lynn–Shacham的简称，[see](https://www.huaweicloud.com/articles/8070926317314981b8e684133ce7ce8b.html)
TLS| |传输层安全性协议（英语：Transport Layer Security，缩写：TLS）
SSL| |安全套接层（英语：Secure Sockets Layer，缩写：SSL）
CA| |数字证书认证机构（英语：Certificate Authority，缩写为CA）
digital certificate|数字证书|
x509|x509 | 是密码学里公钥证书的格式标准
