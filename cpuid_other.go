// Minio Cloud Storage, (C) 2016 Minio, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// +build ppc64 ppc64le mips mipsle mips64 mips64le s390x wasm gccgo

package sha256

func cpuid(op uint32) (eax, ebx, ecx, edx uint32) {
	return 0, 0, 0, 0
}

func cpuidex(op, op2 uint32) (eax, ebx, ecx, edx uint32) {
	return 0, 0, 0, 0
}

func xgetbv(index uint32) (eax, edx uint32) {
	return 0, 0
}

func haveArmSha() bool {
	return false
}
