# Build our rust library
cd lib/obao
cargo build --release

# Move the binary to the correct location
STATIC_TARGET=libobao.a
DYNAMIC_TARGET=libobao.so

# If the static target exists, move it up to the lib directory
if [ -f $STATIC_TARGET ]; then
    mv $STATIC_TARGET ../../lib
fi
# If the dynamic target exists, move it up to the lib directory
if [ -f $DYNAMIC_TARGET ]; then
    mv $DYNAMIC_TARGET ../../lib
    echo "Remember to add this directory to your LD_LIBRARY_PATH:"
fi
