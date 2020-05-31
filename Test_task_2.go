package main

import ("fmt"
        "database/sql"
        _ "github.com/go-sql-driver/mysql"

        "os"
      )

type first_query struct{
  title string
  first_name string
  last_name string
  salary float64
}

type second_query struct{
  department string
  title string
  first_name string
  last_name string
  hire_date string
  work_experience int
}

type third_query struct{
  department string
  count_emp string
  sum_of_sal float64
}

func main(){
  var login string
  var password string
  fmt.Print("Login: ")
  fmt.Fscan(os.Stdin, &login)
  fmt.Print("Password: ")
  fmt.Fscan(os.Stdin, &password)
  db, err := sql.Open("mysql", login + ":" + password + "@/employee")

  if err != nil {
    panic(err)
    }
  defer db.Close()

  var n int
  fmt.Println("What's query get do you want?(1, 2 or 3): ")
  fmt.Println("\n" + `1) Find all current managers of each department and display his/her title,
  first name, last name, current salary.`)
  fmt.Println("\n" + `2) Find all employees (department, title, first name, last name, hire date,
  how many years they have been working) to congratulate them on their hire anniversary this month`)
  fmt.Println("\n3) Find all departments, their current employee count, their current sum salary.\n")

  fmt.Fscan(os.Stdin, &n)

  switch n {
  case 1: First_query(db)
  case 2: Second_query(db)
  case 3: Third_query(db)
  default: return
}
}
func First_query(db *sql.DB){
  query :=
  `SELECT  t.title, e.first_name, e.last_name, s.salary
    FROM employees.dept_manager d
        JOIN employees.employees e USING(emp_no)
        JOIN employees.titles t USING(emp_no)
        JOIN employees.salaries s USING(emp_no)
            WHERE d.to_date >= NOW()
            AND t.to_date >= NOW()
            AND s.to_date >= NOW();`

  rows, err := db.Query(query)
  if err != nil {
    panic(err)
  }
  defer rows.Close()

  first_queries := []first_query{}
  for rows.Next(){
    fq := first_query{}
    err := rows.Scan(&fq.title, &fq.first_name, &fq.last_name, &fq.salary)
    if err != nil{
        fmt.Println(err)
        continue
    }
    first_queries = append(first_queries, fq)
  }

   for _, fq := range first_queries{
       fmt.Println(fq.title, fq.first_name, fq.last_name, fq.salary)
   }
}

func Second_query(db *sql.DB){
  query :=
  `SELECT d.dept_name, t.title, e.first_name, e.last_name, e.hire_date,
        EXTRACT(YEAR FROM NOW()) - EXTRACT(YEAR FROM e.hire_date) as work_experience
    FROM departments d
        JOIN dept_emp d_e USING(dept_no)
        JOIN employees e USING(emp_no)
        JOIN titles t USING(emp_no)
            WHERE d_e.to_date >= NOW()
            AND t.to_date >= NOW();`

  rows, err := db.Query(query)
  if err != nil {
    panic(err)
  }
  defer rows.Close()

  second_queries := []second_query{}
  for rows.Next(){
    fq := second_query{}
    err := rows.Scan(&fq.title, &fq.first_name, &fq.last_name, &fq.department, &fq.work_experience)
    if err != nil{
        fmt.Println(err)
        continue
    }
    second_queries = append(second_queries, fq)
  }

   for _, fq := range second_queries{
       fmt.Println(&fq.title, &fq.first_name, &fq.last_name, &fq.department, &fq.work_experience)
   }
}

func Third_query(db *sql.DB){
  query :=
  `SELECT d.dept_name, COUNT(d_e.emp_no), SUM(s.salary)
    FROM departments d
        JOIN dept_emp d_e USING(dept_no)
        JOIN salaries s USING(emp_no)
            WHERE d_e.to_date >= NOW()
            AND s.to_date >= NOW()
                GROUP BY d.dept_name;`

  rows, err := db.Query(query)
  if err != nil {
    panic(err)
  }
  defer rows.Close()

  third_queries := []third_query{}
  for rows.Next(){
    fq := third_query{}
    err := rows.Scan(&fq.department, &fq.count_emp, &fq.sum_of_sal)
    if err != nil{
        fmt.Println(err)
        continue
    }
    third_queries = append(third_queries, fq)
  }

   for _, fq := range third_queries{
       fmt.Println(&fq.department, &fq.count_emp, &fq.sum_of_sal)
   }
}
