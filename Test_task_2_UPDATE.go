package main

import ("fmt"
        "database/sql"
        _ "github.com/go-sql-driver/mysql"
        "os"
      )

type Idata_of_struct interface{
  Create() Idata_of_struct
  list_of_attributes(sql.Rows) error
  get_query()string
  to_Print()
}

type first_query struct{
  title string
  first_name string
  last_name string
  salary int
}
func(*first_query) Create()Idata_of_struct{
  return &first_query{}
}
func(fq *first_query) list_of_attributes(rows sql.Rows)error{
  return rows.Scan(&fq.title, &fq.first_name, &fq.last_name, &fq.salary)
}
func(*first_query) get_query()string{
  return `SELECT  t.title, e.first_name, e.last_name, s.salary
            FROM employees.dept_manager d
                JOIN employees.employees e USING(emp_no)
                JOIN employees.titles t USING(emp_no)
                JOIN employees.salaries s USING(emp_no)
                  WHERE d.to_date >= NOW()
                  AND t.to_date >= NOW()
                  AND s.to_date >= NOW();`
}
func(fq *first_query) to_Print(){
  fmt.Printf("%-13v|",fq.title)
  fmt.Printf("%-13v|",fq.first_name)
  fmt.Printf("%-13v|",fq.last_name)
  fmt.Printf("%-13v|\n",fq.salary)
}


type second_query struct{
  department string
  title string
  first_name string
  last_name string
  hire_date string
  work_experience int
}
func(*second_query) Create()Idata_of_struct{
  return &second_query{}
}
func(*second_query) get_query()string{
  return `SELECT d.dept_name, t.title, e.first_name, e.last_name, e.hire_date,
              EXTRACT(YEAR FROM NOW()) - EXTRACT(YEAR FROM e.hire_date) as work_experience
          FROM employees.departments d
            JOIN employees.dept_emp d_e USING(dept_no)
            JOIN employees.employees e USING(emp_no)
            JOIN employees.titles t USING(emp_no)
              WHERE d_e.to_date >= NOW()
              AND t.to_date >= NOW();`
}
func(sq *second_query) list_of_attributes(rows sql.Rows)error{
  return rows.Scan(&sq.department, &sq.title, &sq.first_name, &sq.last_name, &sq.hire_date, &sq.work_experience)
}
func(sq *second_query) to_Print(){
  fmt.Printf("%-18v|",sq.department)
  fmt.Printf("%-18v|",sq.title)
  fmt.Printf("%-13v|",sq.first_name)
  fmt.Printf("%-13v|",sq.last_name)
  fmt.Printf("%-13v|",sq.hire_date)
  fmt.Printf("%-3v|\n",sq.work_experience)
}


type third_query struct{
  department string
  count_emp int
  sum_of_sal int
}
func(*third_query) Create()Idata_of_struct{
  return &third_query{}
}
func(*third_query) get_query()string{
  return `SELECT d.dept_name, COUNT(d_e.emp_no), SUM(s.salary)
            FROM employees.departments d
              JOIN employees.dept_emp d_e USING(dept_no)
              JOIN employees.salaries s USING(emp_no)
                WHERE d_e.to_date >= NOW()
                AND s.to_date >= NOW()
                    GROUP BY d.dept_name;`
}
func(tq *third_query) list_of_attributes(rows sql.Rows)error{
  return rows.Scan(&tq.department, &tq.count_emp, &tq.sum_of_sal)
}
func(tq *third_query) to_Print(){
  fmt.Printf("%-18v|",tq.department)
  fmt.Printf("%-10v|",tq.count_emp)
  fmt.Printf("%-13v|\n",tq.sum_of_sal)
}


func main(){
  db, err := sql.Open("mysql", login_password())
  if err != nil {
    panic(err)
  }
  defer db.Close()

  for true{
    res := choose_query()
    if res == nil{
      fmt.Println("Bye!")
      return
    }

    rows, err := db.Query(res.get_query())
    if err != nil {
      panic(err)
    }
    defer rows.Close()

    queries := parse_query(rows, res)
    for i,p := range(queries){
      p.to_Print()
      if i >= 10 {
        break
      }
    }
  }
}

func login_password() string{
  var login string
  var password string
  fmt.Print("Login: ")
  fmt.Fscan(os.Stdin, &login)
  fmt.Print("Password: ")
  fmt.Fscan(os.Stdin, &password)
  return login + ":" + password + "@/employees"
}
func choose_query()Idata_of_struct{
  fmt.Println("\nWhat's query get do you want?(1, 2 or 3): ")
  fmt.Println("\n" + `1) Find all current managers of each department and display his/her title,
  first name, last name, current salary.`)
  fmt.Println("\n" + `2) Find all employees (department, title, first name, last name, hire date,
  how many years they have been working) to congratulate them on their hire anniversary this month`)
  fmt.Println("\n3) Find all departments, their current employee count, their current sum salary.\n")

  var n int
  fmt.Fscan(os.Stdin, &n)
  switch n {
    case 1: return &first_query{}
    case 2: return &second_query{}
    case 3: return &third_query{}
    default: return nil
  }
}
func parse_query(rows *sql.Rows, dfq Idata_of_struct)[]Idata_of_struct{
  queries := []Idata_of_struct{}
  for rows.Next(){
    nq := dfq.Create()
    err := nq.list_of_attributes(*rows)
    if err != nil{
        fmt.Println(err)
        continue
    }
    queries = append(queries, nq)
  }
  return queries
}
