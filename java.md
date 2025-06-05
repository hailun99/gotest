# java
## java基础语法
public static void main(String[] args)

访问修饰符 关键字 返回值类型 方法名(参数类型 参数名)

## 注释
单行注释 //

多行注释 /* */

文档注释 /**  */ 

## java对象和类
```
类（Class）:

定义对象的蓝图，包括属性和方法。

示例：public class Car { ... }

对象（Object）:

类的实例，具有状态和行为。

示例：Car myCar = new Car();

继承（Inheritance）：

一个类可以继承另一个类的属性和方法。

示例：public class Dog extends Animal { ... }

封装（Encapsulation）：

将对象的状态（字段）私有化，通过公共方法访问。

示例：
private String name; 
public String getName() { return name; }

多态（Polymorphism）：

对象可以表现为多种形态，主要通过方法重载和方法重写实现。

示例：
方法重载：public int add(int a, int b) { ... } 和 public double add(double a, double b) { ... }

方法重写：@Override public void makeSound() { System.out.println("Meow"); }

抽象（Abstraction）：

使用抽象类和接口来定义必须实现的方法，不提供具体实现。

示例：

抽象类：public abstract class Shape { abstract void draw(); }

接口：public interface Animal { void eat(); }

接口（Interface）：

定义类必须实现的方法，支持多重继承。

示例：public interface Drivable { void drive(); }

方法（Method）：

定义类的行为，包含在类中的函数。

示例：public void displayInfo() { System.out.println("Info"); }

方法重载（Method Overloading）：

同一个类中可以有多个同名的方法，但参数不同。
```

## 基本数据类型
```
两大类型 : 内置数据类型   引用数据类型

byte： byte 数据类型是8位、有符号的，以二进制补码表示的整数

short： hort 数据类型是 16 位、有符号的以二进制补码表示的整数

int： int 数据类型是32位、有符号的以二进制补码表示的整数；

long： long 数据类型是 64 位、有符号的以二进制补码表示的整数；

float： float 数据类型是单精度、32位、符合IEEE 754标准的浮点数

double： 浮点数的默认类型为 double 类型；

boolean： true 和 false

char：char 类型是一个单一的 16 位 Unicode 字符
```

## 变量类型
```
局部变量（Local Variables）：

局部变量是在方法、构造函数或块内部声明的变量

数据类型 变量名称 = 数据值;

int localVar = 10; // 局部变量

实例变量（Instance Variables）： 

实例变量是在类中声明，但在方法、构造函数或块之外

int instanceVar; // 实例变量

静态变量或类变量（Class Variables）：

类变量是在类中用 static 关键字声明的变量，它们属于类而不是实例

private static int staticVar; // 类变量

参数变量（Parameters）： 

参数是方法或构造函数声明中的变量，用于接收调用该方法或构造函数时传递的值，参数变量的作用域只限于方法内部
```

## 变量命名规则
```
局部变量：使用驼峰命名法。应该以小写字母开头。变量名应该是描述性的，能够清晰地表示其用途。

实例变量（成员变量）：使用驼峰命名法。应该以小写字母开头。

静态变量（类变量）：使用驼峰命名法，应该以小写字母开头。通常也可以使用大写蛇形命名法，全大写字母，单词之间用下划线分隔。

常量：使用全大写字母，单词之间用下划线分隔。常量通常使用 final 修饰。

参数：使用驼峰命名法。应该以小写字母开头。

类名：使用驼峰命名法。应该以大写字母开头。
```

## 循环
```
**while：**

in x = 10;
while (x < 20) {
    System.out.print("value of x:" + x);
    x++;
    System.out.print("\n");
}

**do...while：**

int x = 10;

do{
     System.out.print("value of x : " + x );
     x++;
     System.out.print("\n");
}while( x < 20 );
```

## ctrl + alt + v 补齐左边
## 选中内容 ctrl + alt + v 选择循环语句
## alt + 回车 // 查看是什么错误
## alt + 鼠标左键 // 竖着选中
## Ctrl + alt + m //  抽取方法
## ctrl + R // 替换

## 面向对象
类与对象
```
public class 类名 {
    属性
    行为(动作)
}

类名 对象名 = new 类名();

Phone p = new Phone();

**键盘录入**

nextInt(参数) // 随机数
nextDouble() // 随机小数
next() // 读取字符串 
遇到空格，制表符，回车就停止接受，这些符号后面的数据就不会接受了

nextLine() // 接收字符串
可以接收空格，制表符，遇到回车才停止接受数据
```

## String概述
```
public string() // 创建一个空字符串
public string(String s) // 根据传入的字符串，创造字符串对象
public string(char[] chs) // 根据字符数组，创造字符串对象
public string(byte[] chs) // 根据字节数组，创造字符串对象

```
### 字符串比较
```
**new 出来的:**

数字比较大小

字符串比较地址值
String s1 = "hello";
String s2 = "hello";
s1 == s2; // false

**直接赋值:**
直接赋值会形成地址，再次赋值会取得到地址
String s1 = "hello";
String s2 = "hello";
s1 == s2; // true

boolean equals方法(比较字符串)  // 完全一样true, 否则false
boolean equalsIgnoreCase(比较字符串)    // 忽略大小写(一般用在验证码)
```

### 遍历字符串
```
public char charAt(int index): // 根据索引返回字符
public int length(): // 返回此字符串的长度
数组长度: 数组名.length
字符串的长度: 字符串对象.length()

String  substring(int beginIndex, int endIndex): // 截取 返回从beginIndex到endIndex的子串

String replace(旧值，新值) // 替换
```

### String类
```
构造方法
String Builder() // 创建一个可变字符串
String Builder(String str) // 根据字符内容,创建一个可变字符串对象

常见方法
public StringBuilder append(任意类型) // 添加数据，返回本身
public StringBuilder reverse() // 翻转内容
public int length() // 返回长度
public String toString() // 转为字符串

stringStringBuilder sb = new StringBuilder("abc"); // abc
sb.append(123); // abc123

String str = sb.toString(); // 反转为String类型

StringJoiner的成员方法 // 可以指定间隔符， 开始符号 结束符号
public StringJoiner add(参数) // 添加数据，返回本身
public int length() // 返回长度
public String toString() // 转为字符串
```


## 集合
```
数组 长度固定 可以存储基本数据类型、引用数据类型

集合 长度可变 可以存储基本数据类型、基本数据类型则需要**包装类**

ArrayList() // 创建一个集合

ArrayList<String> list = new ArrayList<>(); // 规定为String类型

ArrayList list = new ArrayList(); // 任意类型

成员方法
public boolean add(E e) // 添加数据
public boolean remove(E e) // 删除数据
public E set(E e) // 修改指定索引的元素
public E get(int index) // 获取指定索引的数据
public int size() // 获取集合长度
```

## 面向对象进阶
### static 静态
#### 被static修饰的成员变量，叫做静态类变量
```

特点: 被该类所有对象共享 不属于对象，属于类  随着类的加载而加载，优先于对象存在

调用方法: 类名调用(推荐)  对象名调用

private static String teachername; // 加上static 表示共享这个类 (teachername)

```

#### 被static修饰的成员变量，叫做静态方法
```
特点: 多用在测试类和工具类中    javaBean类中很少会用

调用方法：类名调用(推荐)    对象名调用

总结：静态方法中，只能访问静态

非静态可以访问所有

静态中没有this关键字

```
### 重新认识main方法
```
public static void main(String[] args)

public 被JVM(虚拟机)调用,访问权限足够大

static  被JVM(虚拟机)调用,不用创造对象，直接访问  因为main方法是静态的，所有测试类中其他方法也需要是静态的

void 被JVM(虚拟机)调用，不需要返回JVM返回值

main 方法是程序的入口

String[] args 现在没用
```

### 继承
```
public class 类名 extends 父类名{}
```

#### 访问特点
```
就近原则: 谁离我近，我就用谁

this // 访问本类成员变量
super // 访问父类成员变量

System.out.println(name) 局部位置找

System.out.println(this.name) 本类位置找

System.out.println(super.name) 父类位置找
```

#### 方法重写
```
@Override

父类中没有的就需要重写

this // 局部变量  理解为一个变量,表示当前方法调用的地址值
super // 父类对象
```

#### 多态
```
什么是多态?
对象的多种形态

多态前提条件?
有继承/实现关系
有父类引用指向子类对象
有方法重载

多态是什么?
使用父类型作为参数，可以接收所有子类对象
体现多态的扩张性与便利
```

##### 多态总结
```
调用成员变量的特点: 编译看左边，运行也看左边

调用成员方法的特点: 编译看左边，运行看右边
```

#### 多态的优势与弊端
```
优势: 方法中，使用父类型作为参数，可以接收所有子类对象

弊端: 不能使用子类特有功能

类型转换的几种方式: 自动类型转换，强制类型转换

强制类型转换能解决什么问题: 可以转换为真正的子类类型，从而调用子类独有功能   转换类型与真实对象类型不一致会报错

instanceof // 转换时判断关键字

if (a instanceof b){} // 判断a是不是b类型

if (a instanceof Dog b){} // 先判断a是不是Dog类型，如果是,则强转为Dog类型，转换之后变量名为b

```

Shift + F6 // 批量修改
Ctrl + Shift + u // 批量修改大小写

final // 锁定变量，不能被修改


### 抽象方法
```
抽象方法定义格式: public abstract 返回值类型 方法名(参数列表); // 有抽象方法的类一定时抽象类

抽象类定义格式: public abstract class 类名{}

抽象类的作用: 无法确认方法体，就把方法定义为抽象类，强制让子类按照某种格式重写  抽象方法所在的类，必须时抽象类

继承抽象类: 重写抽象类中所有抽象方法

```

#### 接口
```
pubilc interface 接口名{ } // 关键字

public class 类名 implements 接口名{ } // 实现接口 可以实现多个接口 也可以继承一个类时同时实现多个接口

接口与类之间的关系: 

类和类的关系: 继承关系，只能单继承，不能多继承，但可以多层继承

类和接口的关系: 实现关系，可以单继承，也可以多实现，还可以在继承一个类的同时实现多个接口

接口与接口的关系: 继承关系，可以单继承，也可以多继承
```


#### jdk8后接口添加的方法
```
默认方法:

default // 关键字 作用: 解决接口升级问题

定义格式: 

public default 返回值类型 方法名称(参数列表){} // 格式

public default void show(){} // 范例

注意事项: 默认方法不是抽象类，不需要强制被重写。重写时去掉关键字default

静态方法:
public static 返回值类型 方法名称(参数列表){}

public static void show(){} 

注意事项:

静态方法只能通过接口名调用，不能通过实现类或者对象名调用
```

#### 内部类
```
写在一个类里面的类

匿名内部类的格式:
new 类名或者接口名(){
    重写方法
};

什么是匿名内部类:
隐藏了名字的内部类，可以写在成员位置，也可以写在局部位置

格式的细节:
包含了继承或实现，方法重写，创建对象，
整体就是一个类的子类对象或者接口的实现类对象

使用场景:
方法的参数是接口或者类时，
以接口为例，可以传递这个接口的实现类对象，
如果实现类只要使用一次，那么就可以使用匿名内部类

```


## 集合
### List系列集合:
```
添加的元素是有序、可重复、有索引
```

### Set系列集合:
```
添加的元素是无序、不可重复、无索引
```


### Collection接口
```
punlic boolean add(E e){} // 添加元素
public void clear(){} // 清空集合
public boolean remove(E e) // 删除元素
pubilc boolean contains(Object o) // 判断集合中是否包含某个元素
pubilc boolean isEmpty() // 判断集合是否为空
public int size() // 获取集合长度

```

### Collection遍历
#### 迭代器遍历
```
Interator<E> iterator() // 返回迭代器对象，默认指向当前集合的0索引

boolean hasNext() // 判断是否还有下一个元素
E next() // 返回迭代器指向的元素，并迭代到下一个元素

细节注意:
1.报错NoSuchElementException: 迭代器没有这个元素异常
2.迭代器遍历完毕，指针不会重复
3.循环中只能用一次next方法
4.迭代器遍历，不能用集合的方法进行增删
```

#### 增强for遍历
```
for(元素的数据类型 变量名 : 数组或者集合名) {

}
```

#### Lambda表达式遍历
```
default void forEach(Consumer<? super E> action) // 结合Lambda遍历

完整版的
coll.forEach(new Consumer<String>() {
            @Override
            public void accept(String s) {
                System.out.println(s);
            }
        });


lambda表达式:
// () -> {}
coll.forEach(s -> System.out.println(s) ); // 当只有一行时
```

### list集合
```
void add(int index, E element) // 在指定位置添加元素
E remove(int index) // 删除指定位置的元素
E set(int index, E element) // 修改指定位置的元素
E get(int index) // 获取指定位置的元素

```

#### 数据结构
```
栈 // 后进先出，先进后出
队列 // 先进先出，后进后出
数组 // 查询快，增删慢
链表 // 查询慢，增删快



泛型方法
1.使用类名后面定义的泛型    // 所有的方法都能用
2.在方法申明上定义自己的泛型    // 只有本方法使用

修饰 <类型> 返回值类型 方法名(类型 变量名){}

public <T> void show(T t){}


泛型接口
修饰 interface 接口名(类型){}

public interface list<E>{}


? rextends E: 表示可以传递或者E所有的子类类型
? super E: 表示可以传递或者E所有的父类类型

public static void A(ArrayList<? extends E> list){}
public static void A(ArrayList<? super E> list){}

应用场景:
1.定义类，方法，接口，不确定，就可以定义泛型类，泛型方法，泛型接口
2.类型不确定，但时能知道以后只能传递某个继承体系中的，就可以使用泛型的通配符
泛型的通配符:
    关键点:可以限定类型的范围


ArrayList  Arr: 数组    List: 属于List系列的一员
LinkedList:  Linked: 链表   List: 属于List系列的一员
HashSet      Hash: 哈希     Set: 属于Set系列的一员
TreeSet      Tree: 树      Set: 属于Set系列的一员
LinkedHashMap:  Linked: 链表   HashSet: 父类
```

## 双列集合
```
双列集合一次需要存一对数据,分别为键值对
键不能重复，值可以
键值对是一一对应的，每一个键只能找到自己对应的值
键 + 值这个整体我们称之为"键值对"或者"键值对对像"，在java中叫做"Entry对象"

V put(k kye, V value)   // 添加元素
put细节:
要么添加/覆盖

V remove(Object key)    // 根据键值对删除元素
void clear()    // 删除所有元素
boolean containsKey(Object key) // 判断集合中是否包含指定的键
boolean containsValue(Object value) // 判断集合中是否包含指定的值
boolean isEmpty()   // 判断集合是否为空
int size()  // 返回集合中元素的个数
```

### HashMap的特点
```
HashMap是Map里面的一个实现类
没有额外需要学习的地方，自己使用Map里面的方法就可以了
特点都是由键决定的：无需、不重复、无索引
HashMap跟HashSet底层一样，存储的元素是无序的,都是哈希表结构
```

### TreeMap底层是红黑树结构
```
由键决定特性:不重复、无索引、可排序
可排序:对键进行排序
默认按照键的从小到大进行排序，也可以自己规定键的排序规则

ctrl+p //  查看方法参数
```

### 可变参数
```
可变参数本质上就是一个数组

作用: 在形参中接收多个数据

可变参数的格式:类型...变量名.
int...a

细节:
1.可变参数必须放在形参列表的最后一个参数位置
2.形参列表中可变参数只能有一个
```

### Collections工具类
```
public static <T> boolean addAll(Collection<T> c, T... elements) // 添加多个元素
public static void shuffle(List<?> list) // 打乱集合
```

### 不可变集合的书写格式
```
static<E> List<E> of(E... elements) // 创建一个具有指定元素的List集合对象
static<E> Set<E> of(E... elements) // 创建一个具有指定元素的Set集合对象
static<K, V> Map<K, V> of(E... elements) // 创建一个具有指定元素的Map集合对象

注意: 这个集合不能添加,不能删除,不能修改

Map<String, String> map = Map.cpyof // 创建一个不可变Map集合对象

List: 直接用
Set: 元素不能重复
Map: 元素不能重复、键值对数量最多不能超10个
超过10个用ofEntries方法
```


## Stream流
###
``` 
单列集合  default Stream<E> stream() // Collection中默认的方法
双列集合            无                // 直接使用stream流
数组  public static<T> Stream<T> stream(T[] array) // Arrays工具类中的静态方法
零散数据 public static<T> Stream<T> of(T... values) // Stream接口中的静态方法

注意:
数组必须是引用数据类型的，如果传递基本数据类型，是会把整个数组当一个元素，放到Stream当中
```

### StreamDemo8
```
Stream<T> filter(Predicate<? super T> predicate) // 筛选(过滤)
Stream<T> limit(long maxSize) // 获取前几个元素
Stream<T> skip(long n) // 跳过前几个元素
Stream<T> distinct() // 元素去重，依赖(hashCode和equals方法)
static <T> Stream<T>concat(Stream a, Stream b) // 合并a和b两个流
Stream<T> map(Function<T, R> mapper) // 转换流中的数据类型

注意:
中间方法，返回心得Stream流，原来的Stream流只能使用一次，建议使用链式编程

修改Stream流中的数据，不会影响原来集合或者数组中的数据

Stream流的终结方法
void forEach(Consumer action)   // 遍历
long count()    // 统计个数
toArray()   // 收集流中的数据，放到数组
collect(Collector collector)    // // 收集流中的数据，放到集合
```

### 总结
```
1.Stream流的作用
结合了Lambda表达式,简化集合,数组的操作

2.Stream流的步骤
获取Stream流的对象
使用中间方法处理
使用终结方法处理

3.任何获取Stream流的对象
单列集合  Collection中默认的方法
双列集合  直接使用stream流
数组  Arrays工具类中的静态方法
零散数据  Stream接口中的静态方法

4.参见的方法
中间方法 filter，limit，skip，distinct，concat，map
终结方法 forEach，count，coolect
```

## 方法
### 引用成员方法
```
格式:对象::成员方法
其他类: 其他类对象::方法名
本类:  this::成员方法
父类； super::成员方法
```


## 异常
```
Error: 系统级别错误(属于严重问题)
Exception: 程序员出现问题(属于问题)
运行时异常:RuntimeException及子类,编译阶段不会出现异常提醒.运行时出现的异常(如: 数组索引越界异常)
编译时异常:编译阶段会出现异常提醒.编译时出现的异常(如: 创建文件时,文件已存在)
```

### 抛出异常
```
throw new RuntimeException("异常信息"); // 抛出异常
```

### 自己处理(捕获异常)
```
格式:
try {
    可以出现异常的代码
} catch(异常类名 变量名) {
    异常的处理代码;
}

当代码出现异常时,可以让程序继续往下执行

### throwable的成员
public String getNessage()  // 返回hrowable的描述信息字符串
public String toString()    // 返回此可抛出的简单描述
public void printStackTrace()   // 把异常的错误学习输出在控制台
```

### 抛出处理
```
throws // 写在方法定义处,表示声明一个异常告诉调用者,使用本方法可能会出现那些异常
public void 方法()throws 异常类名1,异常类名2...{
    ...
}


throw 写在方法内，结束方法,手动抛出异常对象，交给调用者，方法中下面的代码不再执行
public void 方法()throw {
    throw new NullPointerException();
}
```


## IO流
```
什么是IO流
存储和读取数据的解决方案
I: input
O: output
流: 水流一样的数据

输出流: 程序 --> 文件
文件不存在创建

输入流: 文件 --> 程序
不存在，会报错

字节流: 可以操作所有类型的文件
字符流: 只能操作文本文件
```
### FileOutputStream(输入)
```
FileOutputStream书写细节
void write(int b) // 写入一个字节
void write(byte[] b) // 写入字节数组
void write(byte[] b, int off, int len) // 写入字节数组的一部分
```

```
windows: \r\n // 回车 换行
unix: \n // 换行
mac: \r // 回车
```

#### 写法 
```
// 续写开关，默认为false表示关闭，如果为true表示创建对象不会清空文件
        FileOutputStream fos = new FileOutputStream("a.txt", true);
        // 写入数据
        String s = "shfoajshdal";
        byte[] bytes1 = s.getBytes();
        fos.write(bytes1);

        // 写出换行符号
        String warp = "\r\n";
        byte[] bytes2 = warp.getBytes();
        fos.write(bytes2);

        String s2 = "989";
        byte[] bytes3 = s2.getBytes();
        fos.write(bytes3);

        // 释放数据
        fos.close();
```

### FilelnputStream(输出)
```
书写格式(读不到的返回-1)

FileInputStream  fis = new FileInputStream("a.txt");
// 读取数据
int b1 = fis.read();
 System.out.println((char)b1); // char 强转

int b2 = fis.read();
System.out.println(b2); // h 104

```

### 循环读取
```
FileInputStream fis = new FileInputStream("a.txt");
        int b;
        while ((b = fis.read()) != -1){     // 两个左括号
            System.out.print((char)b);
        }

```

### 拷贝文件
```
FileInputStream fis = new FileInputStream("a.txt");
FileOutputStream fos = new FileOutputStream("b.txt");
// 拷贝文件
int b;
        while ((b = fis.read()) != -1){     // 两个左括号
            fos.write(b);
        }
        fos.close();
        fis.close();
```

### 方法
```
public int read() // 读取一个字节
public int read(byte[] b) // 读取字节数组
```

```
FileInputStream fis = new FileInputStream("a.txt");
        byte[] bytes = new byte[2];

        // 正确用法：read(byte[]) 返回实际读取的字节数
        int len1 = fis.read(bytes);
        System.out.println(len1); // 输出实际读取的字节数（1或2）

        // 仅转换实际读取的部分
        String str1 = new String(bytes, 0, len1);
        System.out.println(str1);   // ab

        int len2 = fis.read(bytes);
        System.out.println(len2); // 2
        String str2 = new String(bytes, 0, len2);
        System.out.println(str2); // cd

        int len3 = fis.read(bytes);
        System.out.println(len3); // 1
        String str3 = new String(bytes, 0, len3);
        System.out.println(str3); // e

        fis.close();
```

## 多线程
```
什么是线程:
有了多线程，我们就可以让程序同时做多件事情

多线程的作用:
提高效率

多线程的应用场景:
只要你想要多个事情同时运行就需要用到多个线程
```

并发:在同一时刻，有多个指令在单个cpu上**交替执行**
并行:在同一时刻，有多个指令在单个cpu上**同时执行**

### 实现法式
**继承Thrrad类:**
```
优点: 编程比较简单,可以直接使用Thread类中的方法
缺点: 可以扩展性较差,不能再继承其他的类,**无法获取到多线程的执行结果**
```

**Runnable接口:**
优点： 扩展性强，实现接口同时还可以继承其他的类
缺点： 编程相对复杂，不能直接使用Thread类中的方法，**无法获取到多线程的执行结果**

**Callable接口**和**Future接口**:优点：扩展性强，实现接口同时还可以继承其他的类，**可以获得到多线程的执行结果**
缺点： 编程相对复杂，不能直接使用Thread类中的方法


### 成员方法
```
String getName(); // 获取线程名称

void setName(String name); // 设置线程名称(构造方法也可以设置名字)

static Thread currentThread(); // 获取当前线程的对象

static void sleep(long time); // 线程休眠指定时间

setPriority(int newpriority); // 设置线程优先级

final int getPriority();  // 获取线程优先级

final void setDaemon(boolean on); // 设置守护线程

public static void yield(); // 线程让步

public static void join(); // 插入线程
```

### 线程安全
```
synchronized() // 锁，()里的锁对象一定是唯一的

public void run() {
        while (true){
            // 同步代码块
           synchronized (MyThread.class){ // 在同一个文件夹中,同一个文件一定是问题一的
               if (ticket < 100){
                   try {
                       Thread.sleep(10);
                   } catch (InterruptedException e) {
                       throw new RuntimeException(e);
                   }
                   ticket++;
                   System.out.println(getName() + "正在卖第" + ticket + "张票");
               }else {
                   break;
               }
           }
        }
    }
```

#### 同步方法
修饰 符 synchronized 返回值类型 方法名(方法参数){ ... }

**特点1:** 同步方法是在锁里面所有的代码

**特点2:** 锁对象不能自己指定

    非静态: this

    静态: 当前类的字节码文件对象

StringBuilder //  单线程安全

StringBuffer // 多线程同步安全


### lock锁
```
void lock(); 获取锁
void unlock();  释放锁
```

```
static Lock lock = new ReentrantLock(); // 创建锁对象，所有对象共享同一把锁

finally // 不管怎么样，一定会执行里面的语句 

锁不能嵌套，不然会出现死锁
```

```
   // static表示数据共享
    static int ticket = 0;
    static Lock lock= new ReentrantLock();

    @Override
    public void run() {
        while (true){
            // 同步代码块
//            synchronized (MyThread.class){
            lock .lock(); // 获取锁
            try {
                if (ticket == 100){
                    break;
                }else {
                    Thread.sleep(10);
                    ticket++;
                     System.out.println(getName() + "正在卖第" + ticket + "张票");
                }
//            }
            } catch (InterruptedException e) {
                e.printStackTrace();
            } finally {
                lock.unlock();  // 释放锁

            }
        }
    }
```


### 生产者和消费者
```
viod wait(); // 线程等待
void notify();  // 线程唤醒
void notifyAll(); // 唤醒所有线程


阻塞队列的继承结构(接口)

lterable //  迭代器(可以使用增强for循环)
Collection // 单列集合
Queue // 队列
BlockingQueue // 阻塞队列

实现类
ArrBlockingQueue // 数组阻塞队列(有界，有长度的)
LinkedBlockingQueue // 链表阻塞队列(无界)

put // 底层是已经有锁的
take // 底层是已经有锁的
```

### 案例:
#### 生产者
```
package week10.a14waitandnotify;

import java.util.concurrent.ArrayBlockingQueue;

public class Cook extends Thread{

    ArrayBlockingQueue <String> queue;

    public Cook(ArrayBlockingQueue<String> queue) {
        this.queue = queue;
    }

    @Override
    public void run() {
        while(true){
            try {
                queue.put("面条");
                System.out.println( "厨师放了一碗面条");
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }
    }
}

```

####  消费者
```
package week10.a14waitandnotify;

import java.util.concurrent.ArrayBlockingQueue;

public class Foodie extends Thread{

    ArrayBlockingQueue<String> queue;

    public Foodie(ArrayBlockingQueue<String> queue) {
        this.queue = queue;
    }

    @Override
    public void run() {
        while(true){
            try {
                String food = queue.take();
                System.out.println( "吃掉" + food);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }
    }
}

```

####  测试类
```
package week10.a14waitandnotify;

import java.util.concurrent.ArrayBlockingQueue;

public class ThreadDemo {
    public static void main(String[] args) {
        /*
        * 需求:利用阻塞队列完成生产者和消费者(等待唤醒机制)的代码
        * 细节 :
        *       生产者和消费者必须使用同一个阻塞队列
        *
        * */
        // 创建阻塞队列的对象
        ArrayBlockingQueue queue = new ArrayBlockingQueue(1);

        // 创建线程的对象，并把阻塞队列传递过去
        Cook c = new Cook(queue);
        Foodie f = new Foodie(queue);

        // 开启线程
        c.start();
        f.start();

    }
}
```

### 线程状态
```
新建状态(new) ---> 创建线程对象

就绪绪状态(runnable) ---> start方法

阻塞状态( blocked) ---> 无法获取得对象

等待状态(waiting) ---> 调用wait方法

计时等待(timed_waiting) ---> 调用sleep方法

结束状态(terminated) ---> 线程执行完毕
```

### 线程池
Executors: 线程池的工具类通过调用方法返回不同类型的线程池对象
```
public start ExceptionService newCachedThreadPool() // 创建一个没有上限的线程池
public start ExceptionService newFixedThreadPool(int nThreads) // 创建有上限的线程池
```

主要原理:
1.创建一个池子,池子中是空

2.提交任务时，池子会创建新的线程对象，任务执行完毕，线程归还池子，
下回再次提交任务时，不需要创建新的线程，直接复用已有的线程即可

3.如果提交任务时，池子中没有空闲线程，也无法创建新的线程，任务就会排队等待


#### 自定义线程池(任务拒绝策略)
```
ThreadPoolExecutor.AbortPolicy() // (默认策略):丢弃任务并抛出RejectedExecutionException异常。
ThreadPoolExecutor.DiscardPolicy()// 也是丢弃任务，但是不抛出异常。
ThreadPoolExecutor.DiscardOldestPolicy()// 丢弃队列最前面的任务，然后重新尝试执行任务（重复此过程）
ThreadPoolExecutor.CallerRunsPolicy()// 调用任务的run()方法绕过线程池 ，直接执行

1.创建一个空的池子
2.有任务提交时，线程池会创建线程去执行任务，执行完毕归还线程

不断地提交任务，会有以下三个临界点
1.当核心线程满时，在提交任务就会排队
2.当核心线程满,队伍满时，会创建临时线程
3.当核心线程满，队伍满,临时线程满时，会触发任务拒绝策略.
```

## 反射
```
 // 1.第一种方式
        // 全类名: 包名 + 类名
        // 最常用地
        Class clazz1 = Class.forName("week11.myreflect1.Student");

        // 2.第二种方式
        // 一般更多的是当做参数进行传递
        Class clazz2 = Student.class;

        // 3.第三种方式
        // 当已经有了这个类的对象时，才可以使用
        Student s = new Student();
        Class clazz3 = s.getClass();
```

获取class对象 // Class
构造方法 // Constructor
字段(成员变量) // Field
成员方法 // Method

#### Class类中用于获取构造方法的方法
```
Constructor<?>[] getConstructors() // 返回所有公共构造方法对象的数组
Collection<?>[] getDeclaredConstructors() // 返回所有构造方法对象的数组
Constructor<?> getConstructor(Class<?>... parameterTypes) // 返回单个公共构造方法对象
Constructor<?> getDeclaredConstructor(Class<?>... parameterTypes) // 返回单个构造方法对象
```

#### 反射的作用
```
1.获取任意一个类中的所有信息
2.结合配置文件动态创建对象

获取class字节码文件对象的三种方式
Class.forName ("全类名");
类名.class
对象.getClass();

如何获取构造方法、成员方法、成员变量
get: 获取
Constructor:  构造方法
Field: 成员变量
Method: 方法

set: 设置
Parameter: 参数
Modifers: 修饰符
Declared: 私有的
```


## 动态代理
```
1.为什么需要代理
代理可以无侵入式的给对象增强其他的功能
调用者 --> 代理 ---> 被代理对象

2.代理长什么样
代理里面就是对象要被代理的方法

3.java通过什么来保证代理的样子
通过接口保证  后面的对象和代理需要实现同一个接口
接口就是要被代理的所有方法
```

创建一个代理对象:
```
pubilc start Object newProxyInstance(ClassLoader loader, Class<?>[] interfaces, InvocationHandler h)
参数一: 用于指定哪个类加载器, 去加载生成的代理
参数二: 指定接口,这些接口用于指定生成的代理长什么样，也就是有那些方法
参数三: 用来指定生成的代理对象要干什么事情
```

### 案例
```
package week11.mydynamicproxy1;

import java.lang.reflect.InvocationHandler;
import java.lang.reflect.Method;
import java.lang.reflect.Proxy;

public class ProxyUtil {

    /*
    * pubilc start Object newProxyInstance(ClassLoader loader, Class<?>[] interfaces, InvocationHandler h)
        参数一: 用于指定哪个类加载器, 去加载生成的代理
        参数二: 指定接口,这些接口用于指定生成的代理长什么样，也就是有那些方法
        参数三: 用来指定生成的代理对象要干什么事情
    * */
    public static Star createProxy(BigStar bigStar){

        Star star = (Star) Proxy.newProxyInstance(
                ProxyUtil.class.getClassLoader(), //参数一: 用于指定哪个类加载器, 去加载生成的代理
                new Class[]{Star.class}, //参数二: 指定接口,这些接口用于指定生成的代理长什么样，也就是有那些方法

                // 参数三: 用来指定生成的代理对象要干什么事情
                new InvocationHandler() {
                    @Override
                    public Object invoke(Object proxy, Method method, Object[] args) throws Throwable {
                        /*
                        * 参数一:代理的对象
                        * 参数二:代理对象要调用的方法 sing
                        * 参数三:调用sing方法时,传递的实参
                        * */
                        if ("sing".equals(method.getName())){
                            System.out.println("准备话筒,收钱");
                        }else {
                            System.out.println("准备场地,收钱");
                        }
                        return method.invoke(bigStar,args);
                    }
                }
        );
        return star;
    }
}

```



# javaScript











