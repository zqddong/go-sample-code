 ### 工厂模式

[引用](https://mp.weixin.qq.com/s/MlC6-TDf06LGpF8hxcSV_w)

把一些执行流程明确、但流程细节有差异的业务处理逻辑抽象成公共类库。这时候一组具有相同行为的类都会实现一个固定流程的接口，但是程序里该创建哪个类的实例，提前不知道，只能是根据运行参数动态判断，在加上对于类实例化的过程我们可能需要收敛一下，这样才能保证生产出来的实例能符合我们的预期。

于是乎，聪明的你一定会想到，这时候，我让类库再暴露出一个 NewXXX 之类的方法，这个 NewXXX 方法能够根据条件生产出具体类型的实例返回给业务程序用。如果你能想到这里，恭喜你，这个时候你已经用上工厂模式了

工厂模式一共可以提炼成三类工厂
- 简单工厂
- 工厂方法
- 抽象工厂

### 简单工厂
Go 语言没有构造函数一说，所以一般会定义 NewXXX 函数来初始化相关类。NewXXX 函数返回接口时就是简单工厂模式

简单工厂的优点是，简单，缺点嘛，如果具体产品扩产，就必须修改工厂内部，增加Case，一旦产品过多就会导致简单工厂过于臃肿，为了解决这个问题，才有了下一级别的工厂模式--工厂方法。

### 工厂方法

工厂方法模式（Factory Method Pattern）又叫作多态性工厂模式，指的是定义一个创建对象的接口，但由实现这个接口的工厂类来决定实例化哪个产品类，工厂方法把类的实例化推迟到子类中 进行。

在工厂方法模式中，不再由单一的工厂类生产产品，而是由工厂 类的子类实现具体产品的创建。因此，当增加一个产品时，只需增加一个相应的工厂类的子类, 以解决简单工厂生产太多产品时导致其内部代码臃肿（switch … case分支过多）的问题

工厂方法模式的优点
- 灵活性增强，对于新产品的创建，只需多写一个相应的工厂类。
- 典型的解耦框架。高层模块只需要知道产品的抽象类，无须关心其他实现类，满足迪米特法则、依赖倒置原则和里氏替换原则。

工厂方法模式的缺点
- 类的个数容易过多，增加复杂度。
- 增加了系统的抽象性和理解难度。
- 只能生产一种产品，此弊端可使用抽象工厂模式解决。

无论是简单工厂还是工厂方法都只能生产一种产品，如果工厂需要创建生态里的多个产品，就需要更进一步，使用第三级的工厂模式--抽象工厂。

### 抽象工厂

抽象工厂模式：用于创建一系列相关的或者相互依赖的对象。

为了更清晰地理解工厂方法模式和抽象工厂模式的区别，我们举一个品牌产品生态的例子。

比如智能家居领域多家公司，现在有华为和小米，他们的工厂除了生产我们熟知的手机外，还会生产电视、空调这种家电设备。

假如我们有幸作为他们工厂智能化管理软件的供应商，在软件系统里要对工厂进行抽象，这个时候就不能再用工厂方法这种设计模式了，因为工厂方法只能用来生产一种产品

抽象工厂模式的优点
- 当需要产品族时，抽象工厂可以保证客户端始终只使用同一个产品的产品族。
- 抽象工厂增强了程序的可扩展性，对于新产品族的增加，只需实现一个新的具体工厂即可，不需要对已有代码进行修改，符合开闭原则。

抽象工厂模式的缺点
- 规定了所有可能被创建的产品集合，产品族中扩展新的产品困难，需要修改抽象工厂的接口。
- 增加了系统的抽象性和理解难度。