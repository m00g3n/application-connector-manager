#!/bin/bash

function get_k3d_file_name () {

	local _OS_TYPE=$1
	local _OS_ARCH=$2

	[ "$_OS_TYPE" == "Linux"   ] && [ "$_OS_ARCH" == "x86_64" ] && echo "k3d-linux-amd64"		||
	[ "$_OS_TYPE" == "Linux"   ] && [ "$_OS_ARCH" == "arm64"  ] && echo "k3d-linux-arm64"		||
	[ "$_OS_TYPE" == "Windows" ] && [ "$_OS_ARCH" == "x86_64" ] && echo "k3d-windows-amd64.exe"	||
	[ "$_OS_TYPE" == "Darwin"  ] && [ "$_OS_ARCH" == "x86_64" ] && echo "k3d-darwin-amd64"		||
	[ "$_OS_TYPE" == "Darwin"  ] && [ "$_OS_ARCH" == "arm64"  ] && echo "k3d-darwin-arm64"
}

get_k3d_file_name "$@"
