// NOTE: You could use https://michael-f-bryan.github.io/rust-ffi-guide/cbindgen.html to generate
// this header automatically from your Rust code.  But for now, we'll just write it by hand.

typedef struct ObaoData {
    // Pointers to our obao and Blake3 hash
    const char *obao_data;
    const char *hash_data;
    // How long our obao is
    size_t obao_data_len;
    // How long our hash is (This is 32 bytes)
    size_t hash_data_len;

} ObaoData;

// Reads the file from a path and returns a pointer to a struct containing the obao and hash
ObaoData *obao_data(const char *filepath);
