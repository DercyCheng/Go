package main_test

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"
	"text/template"
)

// https://www.liwenzhou.com/posts/Go/go_template/
// https://books.studygolang.com/gopl-zh/ch4/ch4-06.html
// 解析并渲染模板文件
func TestTmpl(t *testing.T) {
	sayHello := func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("./asset/index.tmpl", "UTF8")
		err = tmpl.Execute(w, "world")
		if err != nil {
			defer log.Fatal(err)
			return
		}
	}
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		defer log.Fatal(err)
		return
	}

}

// 模板语法{{.}}
// 注释，执行时会忽略。可以多行。注释不能嵌套，并且必须紧贴分界符始止
func TestField(t *testing.T) { // 字段替换
	type person struct {
		// 首字母大写导出
		Name string
		Age  string
	}
	temp := template.New("hello")
	temp, _ = temp.Parse("hello {{.Name}}{{-/* a comment */}} {{- .Age}} ")
	p := person{Name: "mary", Age: "18"}
	if err := temp.Execute(os.Stdout, p); err != nil {
		fmt.Println("There was an error:", err.Error())
	}
}

// 串联使用函数
func TestValidation(t *testing.T) {
	tOk := template.New("ok")
	tErr := template.New("error_template")
	defer func() {
		if err := recover(); err != nil {
			log.Printf("run time panic: %v", err)
		}
	}()

	// a valid template, so no panic with Must:
	template.Must(tOk.Parse("/* and a comment */ some static text: {{ .Name }}"))
	fmt.Println("The first one parsed OK.")
	fmt.Println("The next one ought to fail.")
	template.Must(tErr.Parse(" some static text {{ .Name }"))

	ttest := template.New("static test")
	ttest = template.Must(ttest.Parse("ttest "))
	err := ttest.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatal(err)
		return
	}

}
func TestIfelse(t *testing.T) {
	tIfElse := template.New("template test")
	// non empty pipeline following if condition
	tIfElse = template.Must(tIfElse.Parse("if-else demo: {{if `anything`}} Print IF part. {{else}} Print ELSE part.{{end}}\n"))
	tIfElse.Execute(os.Stdout, nil)
}
func TestWE(t *testing.T) {
	twe := template.New("template test")
	twe, _ = twe.Parse("{{with `hello`}}{{.}}{{with `mary`}}{{.}}{{end}}{{end}}!\n")
	twe.Execute(os.Stdout, nil)
}
func TestVariable(t *testing.T) {
	tv := template.New("test")
	tv = template.Must(tv.Parse("{{with $3 := `hello`}}{{$3}}{{end}}!\n"))
	tv_err := tv.Execute(os.Stdout, nil)
	if tv_err != nil {
		log.Fatal(tv_err)
		return
	}

	tv = template.Must(tv.Parse("{{with $x3 := `hola`}}{{$x3}}{{end}}!\n"))
	tvErr := tv.Execute(os.Stdout, nil)
	if tvErr != nil {
		log.Fatal(tvErr)
		return
	}

	tv = template.Must(tv.Parse("{{with $x_1 := `hey`}}{{$x_1}} {{.}} {{$x_1}}{{end}}!\n"))
	tverr := tv.Execute(os.Stdout, nil)
	if tverr != nil {
		log.Fatal(tverr)
		return
	}
}

// range-end 结构格式为{{range pipeline}} T1 {{else}} T0 {{end}}
// pipeline 值为数组,切片,通道,字典
func TestRE(t *testing.T) {
	tr := template.New("test")
	tr = template.Must(tr.Parse("{{range .}}{{.}}{{end}}"))
	err := tr.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
}

// 模板预定义函数
func TestPredefined(t *testing.T) {
	tp := template.New("test")
	tp = template.Must(tp.Parse("{{with $x := `hello`}}{{printf `%s %s` $x `Mary`}}{{end}}!\n"))
	err := tp.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
}
