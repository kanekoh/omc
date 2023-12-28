package cmd

import (
	"testing"
)

func Test_getKubeConfigPath(t *testing.T) {
	type args struct {
		homeDir string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "return default kube config path",
			args: args{
				homeDir: "/home/test",
			},
			want: "/home/test/.kube/config",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getKubeConfigPath(tt.args.homeDir); got != tt.want {
				t.Errorf("getKubeConfigPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getKubeConfigPathWithEnv(t *testing.T) {
	t.Setenv("KUBECONFIG", "/tmp/kubeconfig")

	type args struct {
		homeDir string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "return default kube config path",
			args: args{
				homeDir: "/home/test",
			},
			want: "/tmp/kubeconfig",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getKubeConfigPath(tt.args.homeDir); got != tt.want {
				t.Errorf("getKubeConfigPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
