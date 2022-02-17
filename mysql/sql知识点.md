`<SELECT clause> [<FROM clause>] [<WHERE clause>] [<GROUP BY clause>] [<HAVING clause>] [<ORDER BY clause>] [<LIMIT clause>]`

**SQL1 查找最晚入职员工的所有信息**

>  **排序：** 
>
> select * from 表 order by 列 asc	- 根据 “列” 从小到大排列
> select * from 表 order by 列 desc	- 根据 “列” 从大到小排列
> select * from 表 order by 列1 desc,列2 asc	- 根据 “列1” 从大到小排列，如果相同则按列2从小到大排序

```sql
SELECT * FROM employees
WHERE hire_date = (
    SELECT MAX(hire_date) FROM employees
); # 优

SELECT * FROM employees ORDER BY hire_date DESC LIMIT 1
```

 **SQL2 查找入职员工时间排名倒数第三的员工所有信息** 

> **分页：**
>
> select * from 表 limit 5;	- 前5行
> select * from 表 limit 4,5;	- 从第4行开始的5行
> select * from 表 limit 5 offset 4	- 从第4行开始的5行

```sql
SELECT * 
FROM employees
WHERE hire_date =(
    SELECT hire_date
    FROM employees
    GROUP BY hire_date
    ORDER BY hire_date DESC
    # 从第2行后面开始的1行
    LIMIT 2,1
)
```

 **SQL3 查找当前薪水详情以及部门编号dept_no**

> **连表操作：**
>
> **= ：**`select * form a,b where a.x = b.o` 是JOIN 和 INNER JOIN 的隐式写法
>
> **LEFT OUTER JOIN ：**`select * from a LEFT OUTER JOIN b ON a.x = b.o`  A表所有显示，如果B中无对应关系，则值为null
>
> **inner join** ,永远不会出现Null：`select * from a inner JOIN b ON a.x = b.o   `

```sql
SELECT a.* , b.dept_no 
FROM salaries a,dept_manager b
WHERE a.emp_no = b.emp_no
ORDER BY a.emp_no ASC
# where等值连接不建议使用，会形成笛卡尔积，产生临时表，效率低，是JOIN 和 INNER JOIN 的隐式写法

SELECT a.* , b.dept_no 
FROM salaries AS a
INNER JOIN dept_manager AS b
ON a.emp_no = b.emp_no
ORDER BY a.emp_no ASC
```

**SQL4 查找所有已经分配部门的员工的last_name和first_name以及dept_no**

```sql
SELECT a.last_name,a.first_name,b.dept_no
FROM employees a,dept_emp b
WHERE a.emp_no = b.emp_no
# where等值连接不建议使用，会形成笛卡尔积，产生临时表，效率低，是JOIN 和 INNER JOIN 的隐式写法

SELECT
    e.last_name,
    e.first_name,
    d.dept_no
FROM
    dept_emp AS d
INNER JOIN
    employees AS e
ON
    e.emp_no = d.emp_no;
```

**SQL5 查找所有员工的last_name和first_name以及dept_no，无表示为NULL**

```sql
SELECT
    e.last_name,
    e.first_name,
    d.dept_no
FROM
    dept_emp AS d
RIGHT OUTER JOIN # e表所有显示，如果d中无对应关系，则值为null
    employees AS e
ON
    e.emp_no = d.emp_no;
```

**SQL7 查找薪水记录超过15次的员工号emp_no以及其对应的记录次数t**

> select num from 表 group by num
> select num,nid from 表  where nid > 10 group by num,nid order by nid desc
> select num from 表 group by num having max(id) > 10
> 特别的：group by 必须在where之后，order by之前

```sql
SELECT emp_no,COUNT(emp_no)
FROM salaries
GROUP BY emp_no 
HAVING COUNT(emp_no) > 15

SELECT emp_no,COUNT(emp_no) AS ce # 别名
FROM salaries
GROUP BY emp_no 
HAVING ce > 15
```

**SQL8 找出所有员工当前薪水salary情况**

>    对于distinct与group by的使用：
>    1. 当对系统的性能高并且数据量大时使用group by
>    2. 当对系统的性能不高时或者使用数据量少时两者皆可
>    3. 尽量使用group by
>    4.  group by 只会返回第一行数据 

```sql
SELECT DISTINCT salary
FROM salaries
ORDER BY salary DESC


SELECT salary
FROM salaries
GROUP BY salary
ORDER BY salary DESC
```

 **SQL10 获取所有非manager的员工emp_no**      

> MYSQL数据库中, is 和 = 用法
>
> is 仅用在 is null 或 is not null下使用
>
> = 用在两种情况下：判断值是否相等 where id = 1；二是用于赋值set id = 1  

```sql
SELECT a.emp_no
FROM employees AS a
LEFT OUTER JOIN dept_manager AS b
ON a.emp_no = b.emp_no
WHERE b.emp_no is null
```

**SQL11 获取所有员工当前的manager**

获取所有的员工和员工对应的经理，如果员工本身是经理的话则不显示

```sql
# SELECT a.emp_no,b.emp_no managet
# FROM dept_emp AS a
# INNER JOIN dept_manager AS b
# ON a.dept_no = b.dept_no
# WHERE a.emp_no != b.emp_no

SELECT a.emp_no, b.emp_no AS manager_no
FROM dept_emp AS a
LEFT JOIN dept_manager AS b
ON a.dept_no = b.dept_no
WHERE a.emp_no <> b.emp_no
```

**SQL12 获取每个部门中当前员工薪水最高的相关信息**

```sql
SELECT t1.dept_no,t2.emp_no,t1.maxs
FROM
    (# 每个部门最高的薪水
    SELECT a.dept_no,max(b.salary) maxs
    FROM dept_emp AS a
    INNER JOIN salaries AS b
    ON a.emp_no = b.emp_no
    GROUP BY (a.dept_no)
    ) AS t1
INNER JOIN
    (
    SELECT d.emp_no,d.dept_no,s.salary
    FROM dept_emp d
    INNER JOIN salaries s
    ON d.emp_no = s.emp_no
    ) AS t2
ON t1.dept_no = t2.dept_no and t1.maxs = t2.salary
ORDER BY t1.dept_no ASC
```

**SQL15查找employees表emp_no与last_name的员工信息**

```sql
SELECT *
FROM employees
WHERE emp_no % 2 = 1
AND last_name != 'Mary'
# AND last_name <> 'Mary'
ORDER BY hire_date DESC
```

**SQL16 统计出当前各个title类型对应的员工当前薪水对应的平均工资**

请你统计出各个title类型对应的员工薪水对应的平均工资avg。结果给出title以及平均工资avg，并且以avg升序排序 

```sql
SELECT t.title title,AVG(s.salary)
FROM titles AS t
INNER JOIN salaries AS s
ON t.emp_no = s.emp_no
GROUP BY t.title
```

**SQL17 获取当前薪水第二多的员工的emp_no以及其对应的薪水salary**

```sql
SELECT emp_no,salary
FROM salaries
WHERE salary = (
    SELECT salary
    FROM salaries
    ORDER BY salary DESC
    LIMIT 1,1
)
```

**SQL18 获取当前薪水第二多的员工的emp_no以及其对应的薪水salary**        

查找薪水排名第二多的员工编号emp_no、薪水salary、last_name以及first_name，**不能使用order by完成**   

**sqlite通过的代码不一定能通过mysql，因为SQL语法规定，使用聚合函数时，select子句中一般只能存在以下三种元素：常数、聚合函数，group by 指定的列名。如果使用非group by的列名，sqlite的结果和mysql 可能不一样** 

https://blog.csdn.net/john_hongming/article/details/42742965

> SELECT e.emp_no,s.salary,e.last_name,e.first_name
> FROM employees AS e
> INNER JOIN salaries AS s
> ON s.emp_no = e.emp_no
> WHERE salary = (
>     SELECT MAX(salary)
>     FROM salaries
>     WHERE salary <> (
>         SELECT MAX(salary)
>         FROM salaries
>     )
> )

> ```sql
> SELECT e.emp_no,s.salary,e.last_name,e.first_name
> FROM employees AS e
> INNER JOIN salaries AS s
> ON s.emp_no = e.emp_no
> WHERE salary = (
>     # 查找第二多的薪水
>     SELECT s1.salary
>     FROM 
>     salaries s1
>     JOIN
>     # 表的自连接，薪水和多于或等于其薪水的所有组合
>     salaries s2 ON s1.salary <= s2.salary 
>     GROUP BY s1.salary
>     # 选择第二个
>     HAVING COUNT(DISTINCT s2.salary) = 2
> )
> ```
>
> **看！！！！！！！！！！！**

**SQL19 查找所有员工的last_name和first_name以及对应的dept_name**  

查找所有员工的last_name和first_name以及对应的dept_name，也包括暂时没有分配部门的员工

```sql
SELECT
    e.last_name,
    e.first_name,
    d.dept_name
FROM
    (
        SELECT a.dept_name,b.emp_no
        FROM departments AS a
        INNER JOIN dept_emp AS b
        ON a.dept_no = b.dept_no
    ) AS d
RIGHT OUTER JOIN
    employees AS e
ON
    e.emp_no = d.emp_no;
    
    
SELECT
    emp.last_name,
    emp.first_name,
    dep.dept_name
FROM
    employees AS emp
LEFT OUTER JOIN
    dept_emp AS de
ON
    emp.emp_no = de.emp_no
LEFT OUTER JOIN
    departments AS dep
ON
    dep.dept_no = de.dept_no
```

**SQL21 查找在职员工自入职以来的薪水涨幅情况**       

查找在职员工自入职以来的薪水涨幅情况，给出在职员工编号emp_no以及其对应的薪水涨幅growth，并按照growth进行升序 

```sql
# 找growth
SELECT t1.emp_no,t1.afteri - t2.beforei AS growth
FROM(
    # 找当前薪水
    SELECT e.emp_no,s.salary AS afteri
    FROM salaries AS s
    INNER JOIN employees AS e
    ON e.emp_no = s.emp_no
    WHERE s.to_date = '9999-01-01') t1 
JOIN
    # 找入职的薪水
    (SELECT e.emp_no,s.salary AS beforei
    FROM salaries AS s
    INNER JOIN employees AS e
    ON e.emp_no = s.emp_no AND e.hire_date = s.from_date)t2
ON t1.emp_no = t2.emp_no
ORDER BY growth
```

**SQL22 统计各个部门的工资记录数**             

统计各个部门的工资记录数，给出部门编码dept_no、部门名称dept_name以及部门在salaries表里面有多少条记录sum，按照dept_no升序排序，

```sql
SELECT
    departments.dept_no,
    departments.dept_name,
    COUNT(salaries.salary)
FROM
    departments
INNER JOIN
    dept_emp
ON
    dept_emp.dept_no = departments.dept_no
INNER JOIN
    salaries
ON
    dept_emp.emp_no = salaries.emp_no
GROUP BY departments.dept_no
ORDER BY dept_no ASC
```

**SQL23 对所有员工的薪水按照salary降序进行1-N的排名**             

对所有员工的薪水按照salary降序进行1-N的排名，要求相同salary并列，且按照emp_no升序排列

>    1、RANK()  
>
> ​     在计算排序时，若存在相同位次，会跳过之后的位次。
>
> ​     例如，有3条排在第1位时，排序为：1，1，1，4······
>
>   2、DENSE_RANK()  
>
> ​     这就是题目中所用到的函数，在计算排序时，若存在相同位次，不会跳过之后的位次。
>
> ​     例如，有3条排在第1位时，排序为：1，1，1，2······
>
> 3、ROW_NUMBER()  
>
> ​     这个函数赋予唯一的连续位次。
>
> ​     例如，有3条排在第1位时，排序为：1，2，3，4······

```sql
SELECT emp_no,salary,dense_rank() over(ORDER BY salary DESC) AS t_rank
FROM salaries
# ORDER BY salary DESC # 可以不加
```

**SQL24 获取所有非manager员工当前的薪水情况**             

获取所有非manager员工薪水情况，给出dept_no、emp_no以及salary 

```sql
# 找非manager
SELECT d.dept_no,e.emp_no,s.salary
FROM dept_emp AS d
INNER JOIN employees AS e
ON e.emp_no = d.emp_no
INNER JOIN salaries AS s
ON s.emp_no = e.emp_no
WHERE e.emp_no NOT IN(
    # 找manager
    SELECT emp_no
    FROM dept_manager
)



# 可以直接连接两张表
# 找非manager
SELECT d.dept_no,d.emp_no,s.salary
FROM dept_emp AS d
INNER JOIN salaries AS s
ON s.emp_no = d.emp_no
WHERE d.emp_no NOT IN(
    # 找manager
    SELECT emp_no
    FROM dept_manager
)
```

**SQL25 获取员工其当前的薪水比其manager当前薪水还高的相关信息**            

获取员工其当前的薪水比其manager当前薪水还高的相关信息，第一列给出员工的emp_no，第二列给出其manager的manager_no，第三列给出该员工当前的薪水emp_salary，第四列给出该员工对应的manager当前的薪水manager_salary 

```sql
SELECT t1.emp_no,t2.emp_no,t1.salary,t2.salary
FROM(
    # 获取员工薪水
    SELECT s.emp_no,de.dept_no,s.salary
    FROM salaries AS s
    JOIN dept_emp AS de
    ON de.emp_no = s.emp_no) AS t1
JOIN(
    # 获取部门经理的薪水
    SELECT dm.dept_no,dm.emp_no,s.salary
    FROM dept_manager AS dm
    JOIN salaries AS s
    ON dm.emp_no = s.emp_no) AS t2
WHERE t1.dept_no = t2.dept_no AND t1.salary > t2.salary
```

**SQL26 汇总各个部门当前员工的title类型的分配数目**             

汇总各个部门当前员工的title类型的分配数目，即结果给出部门编号dept_no、dept_name、其部门下所有的员工的title以及该类型title对应的数目count，结果按照dept_no升序排序，dept_no一样的再按title升序排序 

```sql
SELECT de.dept_no,d.dept_name,t.title,COUNT(title)
FROM dept_emp AS de
INNER JOIN departments AS d
ON de.dept_no = d.dept_no
INNER JOIN titles AS t
ON t.emp_no = de.emp_no
GROUP BY de.dept_no,t.title
ORDER BY de.dept_no,t.title
```

**SQL29 使用join查询方式找出没有分类的电影id以及名称**  

```sql
SELECT
    film.film_id,
    film.title
FROM
    film
LEFT OUTER JOIN
    film_category
ON
    film_category.film_id = film.film_id
LEFT OUTER JOIN
    category
ON
    category.category_id = film_category.category_id
WHERE category.category_id is NULL
```

**SQL30 使用子查询的方式找出属于Action分类的所有电影对应的title,description**             

```sql
SELECT film.title,film.description
FROM film
WHERE film.film_id IN(
    SELECT film_category.film_id
    FROM film_category
    WHERE film_category.category_id =(
        SELECT category.category_id
        FROM category
        WHERE name = 'Action'
    )
)
```

**SQL32 将employees表的所有员工的last_name和first_name拼接起来作为Name**

请将employees表的所有员工的last_name和first_name拼接起来作为Name，中间以一个空格区分。   **`CONCAT函数`**

```sql
SELECT CONCAT(last_name,' ',first_name)
FROM employees
```

**SQL33 创建一个actor表，包含如下信息**    

> 创建一个actor表，包含如下列信息
>
> | 列表        | 类型        | 是否为NULL | 含义   |
> | ----------- | ----------- | ---------- | ------ |
> | actor_id    | smallint(5) | not null   | 主键id |
> | first_name  | varchar(45) | not null   | 名字   |
> | last_name   | varchar(45) | not null   | 姓氏   |
> | last_update | date        | not null   | 日期   |

```sql
CREATE TABLE actor(
    actor_id SMALLINT(5) NOT NULL PRIMARY KEY,
    first_name VARCHAR(45) NOT NULL,
    last_name VARCHAR(45) NOT NULL,
    last_update DATE NOT NULL
)
# sql
CREATE TABLE actor(
    actor_id SMALLINT(5) NOT NULL PRIMARY KEY,
    first_name VARCHAR(45) NOT NULL,
    last_name VARCHAR(45) NOT NULL,
    last_update TIMESTAMP NOT NULL DEFAULT(datetime('now','localtime'))
)
```

 **SQL34 批量插入数据**     

> insert into 表 (列名,列名...) values (值,值,值...)
> insert into 表 (列名,列名...) values (值,值,值...),(值,值,值...)
> insert into 表 (列名,列名...) select (列名,列名...) from 表 

```sql
INSERT INTO actor(actor_id,first_name,last_name,last_update)
VALUES (1,'PENELOPE','GUINESS','2006-02-15 12:34:33'),
(2,'NICK','WAHLBERG','2006-02-15 12:34:33')
```

**SQL35** **批量插入数据，不使用replace操作** 

对于表actor插入数据,如果数据已经存在，请忽略(不支持使用replace操作)  **INSERT IGNORE INTO**

```mysql
INSERT IGNORE INTO actor(actor_id,first_name,last_name,last_update)
VALUES ('3','ED','CHASE','2006-02-15 12:34:33')
```

**SQL36** **创建一个actor_name表** 

**请你创建一个actor_name表，并且将actor表中的所有first_name以及last_name导入该表.**

> CREATE TABLE actor_name
> SELECT first_name,last_name FROM actor

**SQL37** **对first_name创建唯一索引uniq_idx_firstname** 

> ALTER TABLE actor
> ADD UNIQUE INDEX uniq_idx_firstname(first_name);
> ALTER TABLE actor
> ADD INDEX idx_lastname(last_name)
>
> ~
>
> CREATE UNIQUE INDEX uniq_idx_firstname
> ON actor(first_name);
> CREATE UNIQUE INDEX idx_lastname
> ON actor(last_name)

**SQL38** **针对actor表创建视图actor_name_view** 

针对actor表创建视图actor_name_view，只包含first_name以及last_name两列，并对这两列重新命名，first_name为first_name_v，last_name修改为last_name_v

> CREATE VIEW actor_name_view(first_name_v, last_name_v) AS
> SELECT first_name, last_name FROM actor 

**SQL39** **对salaries表emp_no字段创建索引idx_emp_no** 

查询emp_no为10005，使用强制索引。 

使用 sql 语句查询数据的时候，如果表里有好几个索引，mysql 优化器会自己判断使用哪个索引进行搜索。但是，where 条件比较复杂的时候，优化器使用的索引就不一定是最优索引了。 

https://www.yangdx.com/2020/05/151.html

> SELECT *
> FROM salaries
> **FORCE INDEX** (idx_emp_no)
> WHERE emp_no = 10005

**SQL40** **在last_update后面新增加一列名字为create_date** 

```sql
ALTER TABLE actor
ADD 
COLUMN create_date DATETIME NOT NULL DEFAULT '2020-10-01 00:00:00'
```

**SQL41** **构造一个触发器audit_log** 

在向employees_test表中插入一条数据的时候，触发插入相关的数据到audit中。 

> CREATE TRIGGER audit_log
> AFTER INSERT ON employees_test 
> FOR EACH ROW
> BEGIN
>     insert into audit values(NEW.id,NEW.name);
> END

**SQL42 删除emp_no重复的记录，只保留最小的id对应的记录**

> delete from 表
>
> delete from 表 where id＝1 and name＝'ala'

```sql
# sqlite可以通过
DELETE FROM titles_test
WHERE id not in(
    SELECT min(id)
    FROM titles_test
    GROUP BY emp_no
)
#  MYSQL无法通过，因为查询和更新不能同时，容易发生死锁，构建第三张 
DELETE FROM titles_test
WHERE id not in(
    SELECT *
    FROM(
        SELECT min(id)
        FROM titles_test
        GROUP BY emp_no
    ) AS a
)
```

**SQL43 将所有to_date为9999-01-01的全部更新为NULL** 

将所有to_date为9999-01-01的全部更新为NULL,且 from_date更新为2001-01-01。 

```sql
UPDATE titles_test 
SET to_date = REPLACE(to_date,'9999-01-01',null),from_date = '2001-01-01'
WHERE to_date = '9999-01-01';
```

**SQL44 将id=5以及emp_no=10001的行数据替换成id=5以及emp_no=10005**

```sql
UPDATE titles_test 
SET emp_no = REPLACE(emp_no,10001,10005)
WHERE id = 5
```

**SQL45 将titles_test表名修改为titles_2017** 

> 添加列：alter table 表名 add 列名 类型
> 删除列：alter table 表名 drop column 列名
> 修改列：
>        alter table 表名 modify column 列名 类型;  # 类型
>        alter table 表名 change 原列名 新列名 类型;  # 列名，类型
> 添加主键：
>        alter table 表名 add primary key(列名);
> 删除主键：
>        alter table 表名 drop primary key;
>        alter table 表名 modify 列名 int, drop primary key;
> 添加外键：alter table 从表 add constraint 外键名称
> 删除外键：alter table 表名 drop foreign key 外键名称
> 修改默认值：ALTER TABLE testalter_tbl ALTER i SET DEFAULT 1000;
> 删除默认值：ALTER TABLE testalter_tbl ALTER i DROP DEFAULT;

```sql
ALTER TABLE titles_test RENAME AS titles_2017
```

**SQL46** **在audit表上创建外键约束，其emp_no对应employees_test表的主键id** 

> 在audit表上创建外键约束，其emp_no对应employees_test表的主键id。 
>
> ALTER TABLE audit
> **ADD CONSTRAINT FOREIGN KEY(emp_no)**
> **REFERENCES employees_test(id)**

**SQL48** **将所有获取奖金的员工当前的薪水增加10%** 

```sql
UPDATE salaries SET salary = salary * 1.1
WHERE emp_no IN(
    SELECT emp_bonus.emp_no
    FROM emp_bonus
) and salaries.to_date = '9999-01-01'
```

**SQL50** **将employees表中的所有员工的last_name和first_name通过引号连接起来。** 

```
SELECT CONCAT(last_name,'\'',first_name)
FROM employees

SELECT CONCAT(last_name,"'",first_name)
FROM employees
```

**SQL51** **查找字符串中逗号出现的次数** 

请你统计每个字符串中逗号出现的次数cnt。 

> **mysql内置函数**
>
> CHAR_LENGTH(str)
> 返回值为字符串str 的长度，长度的单位为字符。一个多字节字符算作一个单字符。对于一个包含五个二字节字符集, LENGTH()返回值为 10, 而CHAR_LENGTH()的返回值为5。
>
> CONCAT(str1,str2,...)
> 字符串拼接。如有任何一个参数为NULL ，则返回值为 NULL。
>
> CONCAT_WS(separator,str1,str2,...)
> 字符串拼接（自定义连接符）。CONCAT_WS()不会忽略任何空字符串。 (然而会忽略所有的 NULL）。
>
> CONV(N,from_base,to_base)
> 进制转换。SELECT CONV('a',16,2);表示将 a 由16进制转换为2进制字符串表示
>
> FORMAT(X,D)
> 将数字X的格式以四舍五入的方式保留小数点后 D 位，并将结果以字符串的形式返回。若 D 为 0, 则返回结果不带有小数点，SELECT FORMAT(12332.1,4); 结果为： ‘12,332.1000’
>
> INSERT(str,pos,len,newstr)
> 在str的指定位置插入字符串
> pos：要替换位置起始位置
> len：替换的长度
> newstr：新字符串
> 如果pos超过原字符串长度，则返回原字符串
> 如果len超过原字符串长度，则由新字符串完全替换
>
> INSTR(str,substr)
> 返回字符串 str 中子字符串的第一次出现位置。
>
> LEFT(str,len)
> 返回字符串str 从开始到len位置的子序列字符。
>
> LOWER(str)
> 变小写
>
> UPPER(str)
> 变大写
>
> LTRIM(str)
> 返回字符串 str，移除左边空白
>
> RTRIM(str)
> 返回字符串 str ，结尾空格字符被删去。
>
> SUBSTRING(str,pos,len)
> 获取字符串子序列
>
> LOCATE(substr,str,pos)
> pos起始位置，获取子序列索引位置
>
> REPEAT(str,count)
> 返回一个由重复的字符串str 组成的字符串，字符串str的数目等于count 。
> 若 count <= 0,则返回一个空字符串。若str 或 count 为 NULL，则返回 NULL 。
>
> REPLACE(str,from_str,to_str)
> 返回字符串str 以及所有被字符串to_str替代的字符串from_str 。
>
> REVERSE(str)
> 返回字符串 str ，顺序和字符顺序相反。
>
> RIGHT(str,len)
> 从字符串str 开始，返回从后边最后一个开始len个字符组成的子序列
>
> SPACE(N)
> 返回一个由N空格组成的字符串。
>
> SUBSTRING(str,pos) SUBSTRING(str,pos,len)
> 不带有len 参数的格式从字符串str返回一个子字符串，起始于位置 pos。
> 带有len参数的格式从字符串str返回一个长度同len字符相同的子字符串，起始于位置 pos。

```sql
SELECT id,CHAR_LENGTH(string) - CHAR_LENGTH(REPLACE(string,',',''))
FROM strings
```

**SQL52** **获取employees中的first_name** 

将employees中的first_name，并按照first_name最后两个字母升序进行输出。 

```sql
SELECT first_name
FROM employees
ORDER BY RIGHT(first_name,2)
```

**SQL53** **按照dept_no进行汇总** 

> SELECT dept_no,**GROUP_CONCAT(emp_no)**
> FROM dept_emp
> GROUP BY dept_no

**SQL54** **平均工资** 

查找排除在职(to_date = '9999-01-01' )员工的最大、最小salary之后，其他的在职员工的平均工资avg_salary。 

```sql
SELECT AVG(salary) AS avg_salary FROM salaries 
WHERE to_date = '9999-01-01' 
AND salary NOT IN (SELECT MAX(salary) FROM salaries WHERE to_date = '9999-01-01')
AND salary NOT IN (SELECT MIN(salary) FROM salaries WHERE to_date = '9999-01-01')
```

**SQL55** **分页查询employees表，每5行一页，返回第2页的数据** 

```sql
SELECT * 
FROM employees
LIMIT 5,5
```

**SQL57 使用关键字exists查找未分配具体部门的员工的所有信息**

```sql
SELECT *
FROM employees
WHERE NOT EXISTS(
    SELECT emp_no
    FROM dept_emp
    WHERE employees.emp_no = dept_emp.emp_no
)
```

**SQL59 获取有奖金的员工相关信息。**            

给出emp_no、first_name、last_name、奖金类型btype、对应的当前薪水情况salary以及奖金金额bonus。bonus结果保留一位小数，输出结果按emp_no升序排序。

```sql
SELECT e.emp_no,e.first_name,e.last_name,b.btype,s.salary,(
    CASE b.btype 
    WHEN 1 THEN s.salary * 0.1
    WHEN 2 THEN s.salary * 0.2
    ELSE s.salary * 0.3 END) AS bonus
FROM employees AS e, salaries AS s,emp_bonus AS b
WHERE s.emp_no = e.emp_no AND e.emp_no = b.emp_no AND to_date='9999-01-01'


SELECT e.emp_no,e.first_name,e.last_name,b.btype,s.salary,ROUND((s.salary * b.btype / 10.0),1) AS bonus
FROM employees AS e, salaries AS s,emp_bonus AS b
WHERE s.emp_no = e.emp_no AND e.emp_no = b.emp_no AND to_date='9999-01-01'
```

**SQL60 统计salary的累计和running_total**             

按照salary的累计和running_total，其中running_total为前N个当前( to_date = '9999-01-01')员工的salary累计和，其他以此类推。  

```sql
SELECT emp_no,salary,SUM(salary)OVER(ORDER BY emp_no)AS running_total
FROM salaries
WHERE to_date = '9999-01-01'


SELECT t1.emp_no,t1.salary,SUM(t2.salary)
# 相当于自连接 找小于其工号的所有salary之和
FROM salaries t1,salaries t2
WHERE t1.to_date='9999-01-01' AND t2.to_date='9999-01-01' AND t1.emp_no>=t2.emp_no
GROUP BY t1.emp_no
```

**SQL61 给出employees表中排名为奇数行的first_name**             

在不打乱原序列顺序的情况下，输出：按first_name排升序后，取奇数行的first_name。 如对以上示例数据的first_name排序后的序列为：Anneke、Bezalel、Georgi、Kyoichi。插入顺序为  Georgi Bezalel Kyoichi Anneke，则原序列中的Georgi排名为3，Anneke排名为1，所以按原序列顺序输出Georgi、Anneke。  

```sql
SELECT e1.first_name
FROM employees e1
WHERE(
    # 如果 e1.first_name 是第一位，那 e2.first_name 只有1个
    # 就是 e1.first_name 本身，1%2=1；
    # 如果 e1.first_name 排在第二位，就有它和比它小的2个 
    # 2%2=0，所以不选
    SELECT COUNT(*) 
    FROM employees e2 
    WHERE e1.first_name >= e2.first_name) % 2 = 1 
```

**SQL62 出现三次以上相同积分的情况**  

id为用户主键id，number代表积分情况，让你写一个sql查询，积分表里面出现三次以及三次以上的积分， 若有多个符合条件的number，则按number升序排序输出。 

```sql
SELECT number
FROM grade
GROUP BY number
HAVING COUNT(number) >= 3
ORDER BY number ASC
```

**SQL63 刷题通过的题目排名**             

> 要得到通过题目的排名，比如你通过了8分,你同学a也通过了8分，
>
> 找到大于等于你的成绩，一个9分，一个8分，一个8分，去重复，就一个9，一个8，
>
> count一下总数，第2名，如果有三个同学通过了7个呢，同理，9，8，8，7，7，7 
>
> 后面比这个少的，已经死在了筛选条件，去重，9，8，7，count=3
>
> **WHERE b.number >= a.number**

```sql
SELECT a.id,a.number,(
    SELECT COUNT(DISTINCT b.number)
    FROM passing_number AS b
    WHERE b.number >= a.number
)
FROM passing_number AS a
ORDER BY a.number DESC,a.id ASC
```

**SQL64 找到每个人的任务**         

请你找到每个人的任务情况，并且输出出来，没有任务的也要输出，而且输出结果按照person的id升序排序

```sql
SELECT person.id,person.name,task.content
FROM person
LEFT OUTER JOIN task
ON person.id = task.person_id
ORDER BY person.id ASC
```

**SQL65 异常的邮件概率**             

统计正常用户发送给正常用户邮件失败的概率：有一个邮件(email)表，id为主键， type是枚举类型，枚举成员为(completed，no_completed)，completed代表邮件发送是成功的，no_completed代表邮件是发送失败的。写一个sql查询，每一个日期里面，正常用户发送给正常用户邮件失败的概率是多少，结果保留到小数点后面3位(3位之后的四舍五入)，并且按照日期升序排序

`IF(条件，值1，值2) 如果条件为true，则IF()的返回 （值1），否则返回（值2）  `

```sql
SELECT date,ROUND(SUM(if(type = 'no_completed',1,0))/COUNT(type),3)
FROM email
WHERE send_id IN(
    # 找正常用户
    SELECT id
    FROM user
    WHERE is_blacklist = 0 )
AND receive_id IN(
    # 找正常用户
    SELECT id
    FROM user
    WHERE is_blacklist = 0 )
GROUP BY date
ORDER BY date
```

**SQL66 牛客每个人最近的登录日期(一)**

牛客每天有很多人登录，请你统计一下牛客每个用户最近登录是哪一天。有一个登录(login)记录表， 请你写出一个sql语句查询每个用户最近一天登录的日子，并且按照user_id升序排序 。

```sql
SELECT user_id, MAX(date) AS d
FROM login
GROUP BY user_id
ORDER BY user_id ASC
```

**SQL67 牛客每个人最近的登录日期(二)**             

写出一个sql语句查询每个用户最近一天登录的日子，用户的名字，以及用户用的设备的名字，并且查询结果按照user的name升序排序 

```sql
SELECT user.name, client.name, date
FROM login , user, client
WHERE login.user_id = user.id AND client.id = login.client_id
AND (user.id,date) IN(
    # 找用户 及最近登录时间
    SELECT user_id,MAX(DATE)
    FROM login
    GROUP BY user_id
)
ORDER BY user.name
```

**SQL68 牛客每个人最近的登录日期(三)**             

统计牛客新登录用户的次日成功的留存率. 即第1天登陆之后，第2天再次登陆的概率,保存小数点后面3位(3位之后的四舍五入) 

要统计牛客新登录用户的次日成功的留存率，首先把公式列出来: （第一天登录的新用户并且第二天也登录的用户）/（总用户）即为新登录用户的次日成功的留存率     总用户其实挺好算 

```sql
# 计算总用户
# SELECT COUNT(DISTINCT user_id) FROM login
SELECT ROUND(COUNT(DISTINCT user_id)/(SELECT COUNT(DISTINCT user_id) FROM login),3)
FROM login
WHERE (user_id,date) IN(
    # 找到每个用户第一天登陆的日子加1
    SELECT user_id,DATE_ADD(MIN(date),INTERVAL 1 DAY) 
    FROM login 
    GROUP BY user_id)
```

**SQL69 牛客每个人最近的登录日期(四)**             

写出一个sql语句查询每个日期登录新用户个数，并且查询结果按照日期升序排序 

**看！！！！！！！！！**

```sql
SELECT date,SUM(CASE 
                WHEN (user_id,date) IN (SELECT user_id,MIN(date) FROM login GROUP BY user_id)
                THEN 1
                ELSE 0
                END) AS new
FROM login
GROUP BY date
ORDER BY date
```

**SQL70 牛客每个人最近的登录日期(五)**       

写出一个sql语句查询每个日期新用户的次日留存率，结果保留小数点后面3位数(3位之后的四舍五入)，并且查询结果按照日期升序排序

```sql
SELECT t1.date,IFNULL(ROUND(t2.old/t1.new,3),0)
FROM (
    # 找每天新用户登录的人数 
    SELECT date,SUM(CASE 
                    WHEN (user_id,date) IN (SELECT user_id,MIN(date) FROM login GROUP BY user_id)
                    THEN 1
                    ELSE 0
                    END) AS new
    FROM login
    GROUP BY date) AS t1
LEFT OUTER JOIN(
    # 第二天也登录的用户数量
    SELECT DATE_SUB(date,INTERVAL 1 DAY) AS date,COUNT(DISTINCT user_id) AS old
    FROM login
    WHERE (user_id,date) IN(
        # 找到每个用户第一天登陆的日子加1
        SELECT user_id,DATE_ADD(MIN(date),INTERVAL 1 DAY) 
        FROM login 
        GROUP BY user_id)
    GROUP BY date) AS t2
ON t1.date = t2.date
```

**SQL71 牛客每个人最近的登录日期(六)**  

统计一下牛客每个用户刷题情况，包括: 用户的名字，以及截止到某天，累计总共通过了多少题。 不存在没有登录却刷题的情况，但存在登录了没刷题的情况，会存在提交代码没有通过的情况并记录在刷题表里，通过数目是0。  并且查询结果先按照日期升序排序，再按照姓名升序排序，有登录却没有刷题的那一天的数据不需要输出 

```sql
SELECT u.name,p.date,SUM(number) over(PARTITION BY user_id ORDER BY date)
FROM user AS u
INNER JOIN passing_number AS p
ON u.id = p.user_id
ORDER BY p.date,u.name
```

**SQL72 考试分数(一)**             

请你写一个sql语句查询各个岗位分数的平均数，并且按照分数降序排序，结果保留小数点后面3位(3位之后四舍五入) **`ROUND(AVG(score),3)`**

```sql
SELECT job ,ROUND(AVG(score),3) AS avg
FROM grade
GROUP BY job
ORDER BY AVG(score) DESC


SELECT job ,ROUND(AVG(score),3) AS avg
FROM grade
GROUP BY job
ORDER BY avg DESC
```

**SQL73 考试分数(二)**             

```sql
SELECT id,job,score
FROM grade AS t1
WHERE score >(
    SELECT AVG(score)
    FROM grade AS t2
    WHERE t2.job = t1.job
    GROUP BY job
)
ORDER BY id
```

**SQL74 考试分数(三)**             

找出每个岗位分数排名前2名的用户，得到的结果先按照language的name升序排序，再按照积分降序排序，最后按照grade的id升序排序 

```sql
SELECT g.id,l.name,g.score
FROM (
    # 找排名
    SELECT *,dense_rank() over(PARTITION BY language_id ORDER BY score DESC) AS srank
    FROM grade
)AS g
INNER JOIN language AS l
ON l.id = g.language_id
WHERE g.srank <= 2
ORDER BY l.name,g.score DESC,g.id
```

**SQL75 考试分数(四)**             

写一个sql语句查询各个岗位分数升序排列之后的中位数位置的范围，并且按job升序排序 

```sql
# 先查每一门的数量 如果为基数 则+1除以2 若为偶数 则直接/2 以及 /2+1
SELECT job, floor(CASE 
             WHEN total % 2 = 1 THEN (total + 1)/2 ELSE(total / 2) END),
            floor(CASE 
             WHEN total % 2 = 1 THEN (total + 1)/2 ELSE(total / 2 + 1) END)
FROM(
    SELECT job,COUNT(id) AS total
    FROM grade
    GROUP BY job
) AS g
ORDER BY job




# floor向下取整 ceiling向上取整
SELECT job,floor((COUNT(id)+1)/2), ceiling((COUNT(id)+1)/2)
FROM grade
GROUP BY job
ORDER BY job
```

**SQL76 考试分数(五)**             

写一个sql语句查询各个岗位分数的中位数位置上的所有grade信息，并且按id升序排序 

```sql
SELECT t1.*
FROM (
    # 找排名
    SELECT *,dense_rank() over(PARTITION BY job ORDER BY score DESC) AS srank
    FROM grade
)AS t1
JOIN (
    # 找中位数
    SELECT job,floor((COUNT(id)+1)/2) s, ceiling((COUNT(id)+1)/2) e
    FROM grade
    GROUP BY job
    ORDER BY job) AS t2
ON t2.job = t1.job
AND t1.srank between t2.s and t2.e
ORDER BY t1.id
```

**SQL77 牛客的课程订单分析(一)**             

写出一个sql语句查询在2025-10-15以后状态为购买成功的C++课程或者Java课程或者Python的订单，并且按照order_info的id升序排序  **`DATEDIFF(date,'2025-10-15') >= 1`**

```sql
SELECT *
FROM order_info
WHERE date > '2025-10-15' AND status = 'completed' AND product_name IN ('C++' , 'Java' , 'Python')
ORDER BY id ASC


SELECT *
FROM order_info
WHERE DATEDIFF(date,'2025-10-15') >= 1 AND status = 'completed' AND product_name IN ('C++' , 'Java' , 'Python')
ORDER BY id ASC
```

**SQL78 牛客的课程订单分析(二)**             

写出一个sql语句查询在2025-10-15以后，同一个用户下单2个以及2个以上状态为购买成功的C++课程或Java课程或Python课程的user_id，并且按照user_id升序排序

```sql
SELECT user_id
FROM order_info
WHERE date > '2025-10-15'
AND product_name IN('C++','Python','Java')
AND status = 'completed'
GROUP BY user_id
HAVING COUNT(id) > 1
ORDER BY user_id ASC
```

**SQL79 牛客的课程订单分析(三)**           

> 窗口函数 https://zhuanlan.zhihu.com/p/92654574  
>
> 窗口函数，也叫OLAP函数（Online Anallytical Processing，联机分析处理），可以对数据库数据进行实时分析处理。
>
> 窗口函数的基本语法如下：
>
> 只用sum()会聚合所有的行，sum() over()不会改变行数 
>
> ```text
> <窗口函数> over (partition by <用于分组的列名>
>              order by <用于排序的列名>)
>              
> <窗口函数>的位置，可以放以下两种函数：
> 1） 专用窗口函数，包括后面要讲到的rank, dense_rank, row_number等专用窗口函数。
> 2） 聚合函数，如sum. avg, count, max, min等
> 因为窗口函数是对where或者group by子句处理后的结果进行操作，所以窗口函数原则上只能写在select子句中。
> 
> 窗口函数有以下功能：
> 1）同时具有分组（partition by）和排序（order by）的功能
> 2）不减少原表的行数，所以经常用来在每组内排名
> ```

写出一个sql语句查询在2025-10-15以后，同一个用户下单2个以及2个以上状态为购买成功的C++课程或Java课程或Python课程的订单信息，并且按照order_info的id升序排序 

> ```sql
> SELECT t.id,t.user_id,t.product_name,t.status,t.client_id,t.date
> FROM(
>     SELECT *, COUNT(*) OVER(PARTITION BY user_id) AS num
>     FROM order_info
>     WHERE product_name IN ('C++','Java','Python')
>     AND status = 'completed'
>     AND date >'2025-10-15'
> ) AS t
> WHERE t.num >= 2
> ORDER BY t.id
> ```

```sql
SELECT * 
FROM order_info
WHERE order_info.user_id IN(
    # 查user_id
    SELECT user_id
    FROM order_info
    WHERE date > '2025-10-15'
    AND product_name IN('C++','Python','Java')
    AND status = 'completed'
    GROUP BY user_id
    HAVING COUNT(id) > 1
) AND date > '2025-10-15' 
AND product_name IN('C++','Python','Java') 
AND status = 'completed'
ORDER BY id
```

**SQL80 牛客的课程订单分析(四)**             

写出一个sql语句查询在2025-10-15以后，如果有一个用户下单2个以及2个以上状态为购买成功的C++课程或Java课程或Python课程，那么输出这个用户的user_id，以及满足前面条件的第一次购买成功的C++课程或Java课程或Python课程的日期first_buy_date，以及购买成功的C++课程或Java课程或Python课程的次数cnt，并且输出结果按照user_id升序排序 

```sql
SELECT t.user_id,MIN(t.date),t.num
FROM(
     SELECT *, COUNT(*) OVER(PARTITION BY user_id) AS num
     FROM order_info
     WHERE product_name IN ('C++','Java','Python')
     AND status = 'completed'
     AND date >'2025-10-15'
) AS t
WHERE t.num >= 2
GROUP BY t.user_id
ORDER BY t.user_id
```

**SQL81 牛客的课程订单分析(五)** 

写出一个sql语句查询在2025-10-15以后，如果有一个用户下单2个以及2个以上状态为购买成功的C++课程或Java课程或Python课程，那么输出这个用户的user_id，以及满足前面条件的第一次购买成功的C++课程或Java课程或Python课程的日期first_buy_date，以及满足前面条件的第二次购买成功的C++课程或Java课程或Python课程的日期second_buy_date，以及购买成功的C++课程或Java课程或Python课程的次数cnt，并且输出结果按照user_id升序排序 

> 可考虑使用lead函数，该函数作用是：获取某顺序字段当前记录的下一条记录；具体用法是：lead(date,1,0)over(partition by user_id order by date) as second_date 这样便可以得到第二次购买日期字段 
>
> lead(field, num, defaultvalue) field需要查找的字段，num往后查找的num行的数据，defaultvalue没有符合条件的默认值。 

```sql
SELECT t.user_id,MIN(t.date),second_buy_date,t.num
FROM(
     SELECT *, COUNT(*) OVER(PARTITION BY user_id) AS num,
     LEAD(date,1,0)over(PARTITION BY user_id ORDER BY date)as second_buy_date
     FROM order_info
     WHERE product_name IN ('C++','Java','Python')
     AND status = 'completed'
     AND date >'2025-10-15'
) AS t
WHERE t.num >= 2
GROUP BY t.user_id
ORDER BY t.user_id
```

**SQL82 牛客的课程订单分析(六)**             

写出一个sql语句查询在2025-10-15以后，同一个用户下单2个以及2个以上状态为购买成功的C++课程或Java课程或Python课程的订单id，是否拼团以及客户端名字信息，最后一列如果是非拼团订单，则显示对应客户端名字，如果是拼团订单，则显示NULL，并且按照order_info的id升序排序

```sql
SELECT t.id,t.is_group_buy,t2.name
FROM(
    SELECT *, COUNT(*) OVER(PARTITION BY user_id) AS num
    FROM order_info
    WHERE product_name IN ('C++','Java','Python')
    AND status = 'completed'
    AND date >'2025-10-15'
) AS t
LEFT OUTER JOIN client AS t2
ON t2.id = t.client_id
WHERE t.num >= 2
ORDER BY t.id
```

**SQL83 牛客的课程订单分析(七)**             

写出一个sql语句查询在2025-10-15以后，同一个用户下单2个以及2个以上状态为购买成功的C++课程或Java课程或Python课程的来源信息，第一列是显示的是客户端名字，如果是拼团订单则显示GroupBuy，第二列显示这个客户端(或者是拼团订单)有多少订单，最后结果按照第一列(source)升序排序 

```sql
# 查看 客户端id和次数
SELECT (CASE WHEN client.name is null THEN 'GroupBuy' ELSE client.name END) name,COUNT(t.id)
FROM(
    SELECT *, COUNT(*) OVER(PARTITION BY user_id) AS num
    FROM order_info
    WHERE product_name IN ('C++','Java','Python')
    AND status = 'completed'
    AND date >'2025-10-15'
) AS t
LEFT JOIN client
ON client.id = t.client_id 
WHERE t.num >= 2
GROUP BY name
ORDER BY name




# 查看 客户端id和次数
SELECT ifnull(client.name, 'GroupBuy') AS source,COUNT(t.id)
FROM(
    SELECT *, COUNT(*) OVER(PARTITION BY user_id) AS num
    FROM order_info
    WHERE product_name IN ('C++','Java','Python')
    AND status = 'completed'
    AND date >'2025-10-15'
) AS t
LEFT OUTER JOIN client
ON client.id = t.client_id
WHERE t.num >= 2
GROUP BY client.name
ORDER BY source
```

**SQL84 实习广场投递简历分析(一)**             

写出SQL语句查询在2025年内投递简历的岗位和数量，并且按数量降序排序 

**year(date) = 2025**

```sql
SELECT job, SUM(num) AS cnt
FROM resume_info
WHERE date>='2025-01-01' and date<='2025-12-31'
GROUP BY job
ORDER BY cnt DESC

SELECT job, SUM(num) AS cnt
FROM resume_info
WHERE date BETWEEN '2025-01-01' AND '2025-12-31'
GROUP BY job
ORDER BY cnt DESC

SELECT job, SUM(num) AS cnt
FROM resume_info
WHERE year(date) = 2025
GROUP BY job
ORDER BY cnt DESC
```

**SQL85 实习广场投递简历分析(二)**             

写出SQL语句查询在2025年内投递简历的每个岗位，每一个月内收到简历的数量，并且按先按月份降序排序，再按简历数目降序排序

```sql
SELECT job,LEFT(date,7) AS month, SUM(num) AS cnt
FROM resume_info
WHERE year(date) = 2025
GROUP BY job , month
ORDER BY month DESC,cnt DESC
```

**SQL86 实习广场投递简历分析(三)**   

写出SQL语句查询在2025年投递简历的每个岗位，每一个月内收到简历的数目，和对应的2026年的同一个月同岗位，收到简历的数目，最后的结果先按first_year_mon月份降序，再按job降序排序显示 

```sql
SELECT t1.job,t1.mon,t1.cnt,t2.mon,t2.cnt
FROM(
    # 找2025年
    SELECT job,date,LEFT(date,7) AS mon, SUM(num) AS cnt
    FROM resume_info
    WHERE year(date) = 2025
    GROUP BY job , mon) t1
INNER JOIN(
    # 找2026年
    SELECT job,date,LEFT(date,7) AS mon, SUM(num) AS cnt
    FROM resume_info
    WHERE year(date) = 2026
    GROUP BY job , mon) t2
ON t1.job = t2.job 
AND month(t1.date) = month(t2.date)
ORDER BY t1.mon DESC,t1.job DESC
```

**SQL87 最差是第几名(一)**             

TM小哥和FH小妹在牛客大学若干年后成立了牛客SQL班，班的每个人的综合成绩用A,B,C,D,E表示，90分以上都是A，80~90分都是B，70~80分为C，60~70为D，E为60分以下。假设每个名次最多1个人，比如有2个A，那么必定有1个A是第1名，有1个A是第2名(综合成绩同分也会按照某一门的成绩分先后)。写出一个SQL查询，如果一个学生知道了自己综合成绩以后，最差是排第几名? 结果按照grade升序排序

```sql
SELECT grade,(
    SELECT SUM(number)
    FROM class_grade AS b
    WHERE a.grade >= grade
)
FROM class_grade AS a
ORDER BY grade 
```

```sql
SELECT grade,SUM(number) OVER(ORDER BY grade) AS num
FROM class_grade AS a
ORDER BY grade
```

**SQL88 最差是第几名(二)**             

学生们综合成绩的中位数是什么档位，请你写SQL帮忙查询一下，如果只有1个中位数，输出1个，如果有2个中位数，按grade升序输出 

> 1. 首先求出每个等级的开始名次`left_order`和结束名次`right_order`。 
> 2. 然后求出中位数所在名次`mid`。 
> 3. 如果`mid`在某个等级的开始名次`left_order`和结束名次`right_order`之间，则输出该等级。 

```sql
SELECT DISTINCT grade
FROM 
    # 求出中位数
    (SELECT floor((SUM(number)+1)/2) mid FROM class_grade UNION 
     SELECT ceiling((SUM(number)+1)/2) mid FROM class_grade) t1,
    # 求出每个等级的开始名次和结束名次
    (SELECT grade,
           number,
           SUM(number) over(ORDER BY grade) - number + 1 AS left_order,
           SUM(number) over(ORDER BY grade)          AS right_order
    FROM class_grade
    ORDER BY grade) t2
WHERE t1.mid BETWEEN t2.left_order and t2.right_order
```

**SQL89 获得积分最多的人(一)**             

写一个SQL查找积分增加最高的用户的名字，以及他的总积分是多少(此题数据保证积分最高的用户有且只有1个) 

```sql
SELECT user.name,SUM(grade_info.grade_num) OVER(PARTITION BY grade_info.user_id) AS num
FROM user
INNER JOIN grade_info
ON grade_info.user_id = user.id
ORDER BY num DESC
LIMIT 1
```

```sql
SELECT user.name,SUM(grade_info.grade_num)AS num
FROM user
INNER JOIN grade_info
ON grade_info.user_id = user.id
GROUP BY user.id
ORDER BY num DESC
LIMIT 1
```

**SQL90 获得积分最多的人(二)**             

写一个SQL查找积分增加最高的用户的id(可能有多个)，名字，以及他的总积分是多少，查询结果按照id升序排序 

```sql
# 别名
WITH t AS (
    SELECT user_id,SUM(grade_num) AS grade_sum
    FROM grade_info
    GROUP BY user_id)
SELECT id,name,grade_sum
FROM t
JOIN user u ON u.id = t.user_id
WHERE grade_sum  = (SELECT MAX(grade_sum) FROM t)
```

**SQL91 获得积分最多的人(三)**             

写一个SQL查找积分最高的用户的id，名字，以及他的总积分是多少(可能有多个)，查询结果按照id升序排序 

```sql
# 别名
WITH t AS (
    # 有reduce
    SELECT user_id,SUM(grade_num) AS grade_sum
    FROM (
        SELECT user_id,(CASE type
                        WHEN 'reduce' 
                        THEN -1*grade_num
                        ELSE grade_num
                        END) AS grade_num
        FROM grade_info) AS t1
    GROUP BY user_id)
SELECT id,name,grade_sum
FROM t
JOIN user u ON u.id = t.user_id
WHERE grade_sum  = (SELECT MAX(grade_sum) FROM t)




# 别名
WITH t AS (
    # 有reduce
    SELECT user_id,SUM(grade_num) AS grade_sum
    FROM (
        SELECT user_id, IF(type = 'reduce',-1,1)*grade_num AS grade_num
        FROM grade_info) AS t1
    GROUP BY user_id)
SELECT id,name,grade_sum
FROM t
JOIN user u ON u.id = t.user_id
WHERE grade_sum  = (SELECT MAX(grade_sum) FROM t)
```

**SQL92 商品交易(网易校招笔试真题)**             

查找购买个数超过20,重量小于50的商品，按照商品id升序排序

```sql
SELECT a.id,a.name,a.weight,SUM(b.count) AS cnt
FROM goods AS a
INNER JOIN trans AS b
ON a.id = b.goods_id
GROUP BY b.goods_id
HAVING a.weight < 50 AND cnt > 20
ORDER BY a.id
```

**SQL93 网易云音乐推荐(网易校招笔试真题)** 

编写一个SQL，查询向user_id = 1 的用户，推荐其关注的人喜欢的音乐。 

```sql
SELECT DISTINCT music_name 
FROM music
INNER JOIN music_likes
ON music_likes.music_id = music.id
WHERE music_likes.user_id IN(
    # 查找其关注的人
    SELECT follower_id 
    FROM follow
    WHERE user_id = 1
) AND music_id NOT IN(
    # 查找自己听得音乐
    SELECT music_id
    FROM music_likes
    WHERE user_id  = 1
)
ORDER BY music_id
```

