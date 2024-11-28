package services

import (
	"os/exec"
)

// 初始化 Terraform
func InitTerraform() error {
	cmd := exec.Command("terraform", "init")
	cmd.Dir = "./terraform" // 指定 Terraform 配置目录
	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	println(string(output))
	return nil
}

// 应用 Terraform 配置
func ApplyTerraform() error {
	cmd := exec.Command("terraform", "apply", "-auto-approve")
	cmd.Dir = "./terraform"
	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	println(string(output))
	return nil
}

// 销毁 Terraform 配置
func DestroyTerraform() error {
	cmd := exec.Command("terraform", "destroy", "-auto-approve")
	cmd.Dir = "./terraform"
	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	println(string(output))
	return nil
}
