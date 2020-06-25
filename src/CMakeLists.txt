cmake_minimum_required(VERSION 2.8.4)

SET(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -std=c++11 -pthread")

set(LIBBEAM_SRC 
	./core/ecc.cpp 
	  ./crypto/equihashR_impl.cpp
    ./core/difficulty.cpp
    ./core/ecc.cpp
    ./core/uintBig.cpp
    ./3rdparty/blake/ref/blake2b-ref.c

    ./core/static_members.cpp
    ./pow/beamHash.cpp
    ./utility/common.cpp
	)

include_directories(.)
include_directories(./3rdparty/secp256k1-zkp/include)
include_directories(./3rdparty)
include_directories(./3rdparty/common)

add_library(beam_shared SHARED ${LIBBEAM_SRC})

set_target_properties(beam_shared PROPERTIES OUTPUT_NAME "beampow")

set(LIBRARY_OUTPUT_PATH ${PROJECT_BINARY_DIR}/lib)

install(TARGETS beam_shared
	LIBRARY DESTINATION lib
	ARCHIVE DESTINATION lib)
