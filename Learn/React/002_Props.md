# 组件之间数据传递


### 父组件向子组件传递数据

* 属性传值

```js
// 父组件TodoList
<ul>
	{
		this.state.list.map((item, index) => {
			return(
				<TodoItem content={item}/>
			)
		})
	}
</ul>

// 子组件TodoItem
return <div>{this.props.content}</div>
```


### 子组件向父组件传值

* 调用父组件传递过来的方法


### 生命周期函数

* componentWillMount（组件将要被挂载的时候）
* render（组件渲染）

当props或state改变时，render函数会被重新执行
当父组件的render函数被执行时，子组件的render函数都将被重新执行

* componentDidMount（组件挂载完成）
* componentWillReceiveProps（没有从父组件接收props的组件不会执行这个方法）
* shouldComponentUpdate（返回true更新，返回false更新）
* componentWillUpdate（组件被更新前）
* componentDidUpdate（组件被更新后）
* componentWillUnmount（）


### Props

* PropTypes

数据校验和默认值设置

```js
import PropTypes from 'prop-types';

TodoItem.propTypes = {
	content: PropTypes.string,
	deleteItem: PropTypes.func,
	index: PropTypes.number.isRequired
}

TodoItem.defaultProps = {
	test: 'hello world'
}
```


### State

 
### Forms
