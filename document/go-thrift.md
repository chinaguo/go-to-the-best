http://thrift.apache.org/docs/install/os_x

#### 1.安装boost库

./bootstrap.sh
boost地址 https://www.boost.org/
sudo ./b2 threading=multi address-model=64 variant=release stage install


#### 2.安装libevent
./configure --prefix=/usr/local
make
sudo make install
