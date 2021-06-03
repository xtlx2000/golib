package os

// 关闭计算机
func ShutDownEXE() {
	fmt.Println("关闭主机")
	arg := []string{"-s", "-t", "20"}
	cmd := exec.Command("shutdown", arg...)
	d, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(d))
	return
}

// 重启计算机
func ReShutDownEXE() {
	fmt.Println("重启主机")
	arg := []string{"-r", "-t", "20"}
	cmd := exec.Command("shutdown", arg...)
	d, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(d))
	return
}

// kill调进程  参数---taskkill /im notepad.exe /T /F
// 参数说明：strGameName为需要kill的进程的名字
func KillEXE(strGameName string) bool {
	fmt.Println("kill调进程游戏：", strGameName)
	strGameName = strGameName + ".exe"
	arg := []string{"/im", strGameName}
	cmd := exec.Command("taskkill", arg...)
	if err := cmd.Run(); err != nil {
		fmt.Println("Error: ", err)
	}
	return true
}

//假如你要运行的程序名字为:"autorun.exe"使用命令为
//"reg add HKLM\SOFTWARE\Microsoft\Windows\CurrentVersion\Run /v AUTORUN /t REG_SZ /d C:\autorun.exe /f"
//(不包括引号)其中"C:\autorun.exe"为目标程序的路径.按着这样的命令就可以将你的程序添加到启动项中了
// RunEXE  参数---reg add HKLM\SOFTWARE\Microsoft\Windows\CurrentVersion\Run /v AUTORUN /t REG_SZ /d C:\autorun.exe /f
func RunEXE() {
	fmt.Println("开机启动")
	strEXEName = "C:\\Windows\\System32\\auto.vbs" // 需要启动文件的路径文件
	arg := []string{"add", "HKEY_LOCAL_MACHINE\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Run", "/v", "auto", "/t", "REG_SZ", "/d", strEXEName, "/f"}
	cmd := exec.Command("reg", arg...)
	d, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(d))
	return
}

// 隐藏调进程  参数---start /b notepad.exe
func YinCangEXE(strEXEName string) {
	fmt.Println("隐藏进程")
	cmd := exec.Command("auto.bat")
	d, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(d))
	return
}

// 启动exe
// 参数 strGameName 启动的执行文件的名字；strIPandPort 传递给exe的参数
func CallEXE(strGameName string, strIPandPort string) {
	fmt.Println("CallEXE 开始启动游戏")
	arg := []string{strGameName, strIPandPort}
	fmt.Println("------------", arg)
	strPath := getCurrentPath()
	strPath = strPath + "\\TSTX\\" + strGameName // 路径
	cmd := exec.Command(strPath, arg...)
	if err := cmd.Run(); err != nil {
		fmt.Println("Error: ", err)
		return
	}
	return
}

// 获取当前目录
func getCurrentPath() string {
	s, _ := exec.LookPath(os.Args[0])
	i := strings.LastIndex(s, "\\")
	path := string(s[0 : i+1])
	PathData = path
	return path
}
