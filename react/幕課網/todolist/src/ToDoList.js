import React,{ Component, Fragment} from 'react';
import ToDoItem from './ToDoItem';
import './style.css';
import axios from 'axios';

class ToDoList extends Component{

  constructor(props){ //一定要接收props變數
    super(props);//繼承父類
    this.state = {
      inputValue : '',
      list : []
    }
    this.handleSubmit = this.handleSubmit.bind(this); //bind放於constructor可以使程式較美觀且效率較高
    this.handleDelete = this.handleDelete.bind(this);
    this.handleChange = this.handleChange.bind(this);
  }

  render(){
    return (
      <Fragment> 
        <div>
          <label htmlFor="todolist_input">ToDoList:</label>
          <input 
            id="todolist_input"
            className='input'
            value={this.state.inputValue} type="text"
            onChange={this.handleChange}
            ref={(input) => { this.input_value = input}} //(數據傳入寫法二)(盡量避免使用ref來獲取數值)
          />
          <button type="button" onClick={this.handleSubmit}>提交</button>
        </div>

        <ul  ref={(ul) => { this.ul_value = ul}}>
          {
            this.state.list.map((item , index) =>{
              return (
                  <ToDoItem key={index} handleDelete={this.handleDelete} item={item} index={index}/>
              )
            })
          }
        </ul>
      </Fragment>
    );
  }

  componentDidMount(){ //用於放AJAX事先載入資料時  下方為測試用
  //  axios.get('https://images.unsplash.com/photo-1516550135131-fe3dcb0bedc7?ixlib=rb-0.3.5&ixid=eyJhcHBfaWQiOjEyMDd9&s=621e8231a4e714c2e85f5acbbcc6a730&auto=format&fit=crop&w=1352&q=80').then(
  //   (res)=>{ 
  //     alert('true');
  //   }
  //  ).catch(
  //   ()=>{
  //     alert('false');
  //   }
  //  )
  }

  handleChange(e){
    // this.setState(()=>{          //函式寫法
    //   return {
    //     inputValue : e.target.value
    //   }
    // })
    
    const value = e.target.value    //因為函式寫法為異步執行所以必須先把值保存起來 (數據傳入寫法一)
    const value2 = this.input_value.value //(數據傳入寫法二)

    this.setState(()=>({            //Return 簡寫寫法
      inputValue : value
    }))

    // this.setState({inputValue : e.target.value});   //一般寫法，但es6後不建議
  }

  handleSubmit(){
    this.setState((prevState)=>({  //prevState是紀錄觸發此函術前的State狀態
      list : [...prevState.list , prevState.inputValue],
      inputValue : ''
    }),()=>{ 
        console.log(this.ul_value.querySelectorAll('li').length+'內部'); //setState的內建的Callback函數，setState執行完成後執行，避免異步執行的錯誤
    });

        console.log(this.ul_value.querySelectorAll('li').length+'外部'); //因為setState是異步執行所以不會馬上渲染，導致執行時只會撈到舊的
  }

  handleDelete(index){
    //immutable
    //state不允許直接做任何改變(但是可行的)
    
    this.setState((prevState)=>{
      const list = [...prevState.list];
      list.splice(index, 1); //從index開始刪除1個元素
      return {list:list};  //如果2個名子相同可以簡寫成return {list};
    })
  }


}


export default ToDoList;
