mac下安装thrift
http://thrift.apache.org/docs/install/os_x

#### 1.安装boost库
```js
boost地址: https://www.boost.org/
./bootstrap.sh
sudo ./b2 threading=multi address-model=64 variant=release stage install
```

#### 2.安装libevent

```js
libevent地址: http://libevent.org/
brew install openssl
ln -s /usr/local/Cellar/openssl/1.0.2l/include/openssl/ libevent_home/include/openssl

./configure --prefix=/usr/local
make
sudo make install
```

#### 3.构建thrift

```js
brew install autoconf
brew install libtool
brew install automake
./configure --prefix=/usr/local/ --with-boost=/usr/local --with-libevent=/usr/local
```
