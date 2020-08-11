import React,{ Component, Fragment} from 'react';

class ToDoItem extends Component{

  constructor(props){
    super(props);
    this.handleClick = this.handleClick.bind(this);
  }

  render(){
    return (
      <Fragment>
        {/* <li onClick={this.handleClick}> */}
        <li onClick={()=> this.props.handleDelete(this.props.index)}> {/*子組件要使用父組件的函式時必須由子組件的函示包裹才能使用*/}
          {this.props.item}
        </li>
      </Fragment>
    );
  }

  
  handleClick(){   {/*第2種方法*/}
    this.props.handleDelete(this.props.index);
  }

}



export default ToDoItem;
