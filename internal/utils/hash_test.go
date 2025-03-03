package util

import (
	"reflect"
	"testing"
)

func Test_cipher(t *testing.T) {
	type args struct {
		vec []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{"standard case", args{[]byte{0, 0, 0, 0, 0}}, []byte{13, 14, 15, 16, 17}},
		{"underflow case", args{[]byte{1, 2, 3, 4, 5}}, []byte{14, 16, 18, 20, 22}},
		{"overflown case", args{[]byte{251, 252, 253, 254, 255}}, []byte{8, 10, 12, 14, 16}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := cipher(tt.args.vec)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cipher() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_decipher(t *testing.T) {
	type args struct {
		vec []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{"standard case", args{[]byte{0, 0, 0, 0, 0}}, []byte{243, 242, 241, 240, 239}},
		{"underflow case", args{[]byte{1, 2, 3, 4, 5}}, []byte{244, 244, 244, 244, 244}},
		{"overflown case", args{[]byte{255, 254, 253, 252, 251}}, []byte{242, 240, 238, 236, 234}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decipher(tt.args.vec); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decipher() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bijection(t *testing.T) {
	type args struct {
		vec []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{"standard case", args{[]byte{0, 0, 0, 0, 0}}, []byte{0, 0, 0, 0, 0}},
		{"underflow case", args{[]byte{1, 2, 3, 4, 5}}, []byte{1, 2, 3, 4, 5}},
		{"overflown case", args{[]byte{255, 254, 253, 252, 251}}, []byte{255, 254, 253, 252, 251}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := decipher(cipher(tt.args.vec))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inverse cipher() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToHash(t *testing.T) {
	type args struct {
		id uint64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"id is 0", args{0}, "DQ4PEBES"},
		{"id is 1", args{1}, "Dg4PEBES"},
		{"id is 2", args{2}, "Dw4PEBES"},
		{"id is 3", args{3}, "EA4PEBES"},
		{"id is 4", args{4}, "EQ4PEBES"},
		{"id is 5", args{5}, "Eg4PEBES"},
		{"id is 6", args{6}, "Ew4PEBES"},
		{"id is 7", args{7}, "FA4PEBES"},
		{"id is 8", args{8}, "FQ4PEBES"},
		{"id is 9", args{9}, "Fg4PEBES"},
		{"id is 100", args{100}, "cQ4PEBES"},
		{"id is 1,000", args{1000}, "9REPEBES"},
		{"id is 10,000", args{10_000}, "HTUPEBES"},
		{"id is 100,000", args{100_000}, "rZQQEBES"},
		{"id is 1,000,000", args{1_000_000}, "TVAeEBES"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToHash(tt.args.id)
			if len(got) != LEN_HASH {
				t.Errorf("ToHash() = %v, want a hash with length = %v", err, LEN_HASH)
				return
			}
			if got != tt.want {
				t.Errorf("ToHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToID(t *testing.T) {
	type args struct {
		hash string
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{"id is 0", args{"DQ4PEBES"}, 0},
		{"id is 1", args{"Dg4PEBES"}, 1},
		{"id is 2", args{"Dw4PEBES"}, 2},
		{"id is 3", args{"EA4PEBES"}, 3},
		{"id is 4", args{"EQ4PEBES"}, 4},
		{"id is 5", args{"Eg4PEBES"}, 5},
		{"id is 6", args{"Ew4PEBES"}, 6},
		{"id is 7", args{"FA4PEBES"}, 7},
		{"id is 8", args{"FQ4PEBES"}, 8},
		{"id is 9", args{"Fg4PEBES"}, 9},
		{"id is 100", args{"cQ4PEBES"}, 100},
		{"id is 1,000", args{"9REPEBES"}, 1000},
		{"id is 10,000", args{"HTUPEBES"}, 10_000},
		{"id is 100,000", args{"rZQQEBES"}, 100_000},
		{"id is 1,000,000", args{"TVAeEBES"}, 1_000_000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := ToID(tt.args.hash)
			if got != tt.want {
				t.Errorf("ToID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsValidHash(t *testing.T) {
	type args struct {
		hash string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"normal hash", args{"TVAeEBES"}, true},
		{"to short", args{"foo"}, false},
		{"to long", args{"barbarbarbar"}, false},
		{"bad symbols", args{"????????"}, false},
		{"using space", args{"        "}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidHash(tt.args.hash); got != tt.want {
				t.Errorf("IsValidHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkToHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ToHash(1_000_000)
	}
}

func BenchmarkToID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ToID("TVAeEBES")
	}
}

func BenchmarkCipher(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buf := []byte{0, 1, 2, 3, 4, 5, 6, 7}
		cipher(buf)
		decipher(buf)
	}
}
