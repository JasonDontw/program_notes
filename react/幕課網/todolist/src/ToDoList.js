import React,{ Component, Fragment} from 'react';
import ToDoItem from './ToDoItem';
import './style.css';

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
          />
          <button type="button" onClick={this.handleSubmit}>提交</button>
        </div>

        <ul>
          {
            this.state.list.map((item , index) =>{
              return (
                <div>
                  <ToDoItem key={index} handleDelete={this.handleDelete} item={item} index={index}/>
                </div>
              )
            })
          }
        </ul>
      </Fragment>
    );
  }

  handleChange(e){
    // this.setState(()=>{          //函式寫法
    //   return {
    //     inputValue : e.target.value
    //   }
    // })
    const value = e.target.value    //因為函式寫法為異步執行所以必須先把值保存起來
    this.setState(()=>({            //Return 簡寫寫法
      inputValue : value
    }))

    // this.setState({inputValue : e.target.value});   //一般寫法，但es6後不建議
  }

  handleSubmit(){
    this.setState((prevState)=>({  //prevState是紀錄觸發此函術前的State狀態
      list : [...prevState.list , prevState.inputValue],
      inputValue : ''
    }))
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
