#!/bin/bash
#Maintained by Can Özyurt acozyurt@gmail.com
if command -v varnishd > /dev/null; then
        echo "Varnish installed!"
else
        echo "Installing varnish first... "
        curl -s https://packagecloud.io/install/repositories/varnishcache/varnish60lts/script.deb.sh | sudo bash
		echo -n "Installing varnish... "
        apt install varnish=6.0.6-1 varnish-dev=6.0.6-1 -y &> /dev/null
        echo "done."
fi

echo -n "Updating repository and installing automake... "
apt update &>/dev/null && apt install autoconf automake docutils-common libtool -y &>/dev/null
echo "done."

echo -n "Cloning libvmod-bodyaccess from github... "
git clone https://github.com/aondio/libvmod-bodyaccess.git &> /dev/null
echo "done."

echo -n "Making changes for Varnish 6.0.6 compability... "
sed -i 's/vrt.h/vsb.h/' libvmod-bodyaccess/src/vmod_core.h
sed -i 's/__match_proto__(objiterate_f)//' libvmod-bodyaccess/src/vmod_core.c
sed -i '/^struct/i ssize_t VRB_Iterate(struct req *, objiterate_f *func, void *priv);' libvmod-bodyaccess/src/vmod_core.h
echo "done."

echo -n "Building and installing vmod... "
cd libvmod-bodyaccess && ./autogen.sh &> /dev/null && ./configure &> /dev/null && make &> /dev/null && make install &> /dev/null
echo "done."
echo "Reminder: varnish-modules should be installed manually if needed."