set(CNH_HASHING "cnh_hashing")
set(CMAKE_MACOSX_RPATH 1)

set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -std=c++11 -D_GNU_SOURCE")

add_library(${CNH_HASHING} SHARED src/cnh_hashing.cc)

find_library(CNH_CNCRYPTO cnhcrypto ${CNH_DIR}/src/cnh_crypto)
message("Found CNH ${CNH_CNCRYPTO}")

target_include_directories(${CNH_HASHING} PRIVATE "${CNH_DIR}/contrib/epee/include")
target_include_directories(${CNH_HASHING} PRIVATE "${CNH_DIR}/src")

target_link_libraries(${CNH_HASHING} PRIVATE ${CNH_CNCRYPTO})