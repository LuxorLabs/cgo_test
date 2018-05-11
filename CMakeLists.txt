cmake_minimum_required (VERSION 2.8.11)
project(pool)

get_filename_component(CNH_DIR $ENV{CNH_DIR} ABSOLUTE)
message("Using CryptoNight Heavy source from env ${CNH_DIR}")

add_subdirectory(cgo/cnh_hashing)

add_custom_command(
    OUTPUT build/bin
    COMMAND build/env.sh go get -v ./...
)

add_custom_target(build ALL DEPENDS cnh_hashing ${CMAKE_CURRENT_BINARY_DIR}/build/bin)