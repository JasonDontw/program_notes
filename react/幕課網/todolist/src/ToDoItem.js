import React,{ Component, Fragment} from 'react';
import propTypes from 'prop-types'

class ToDoItem extends Component{

  constructor(props){
    super(props);
    this.handleClick = this.handleClick.bind(this);
  }

  render(){
    return (
      <Fragment>
        {/* <li onClick={this.handleClick}> */} {/*第2種方法*/}
        <li onClick={()=> this.props.handleDelete(this.props.index)}> {/*子組件要使用父組件的函式時必須由子組件的函示包裹才能使用*/}
          {this.props.item}
        </li>
      </Fragment>

    );
  }

  
  handleClick(){   //第2種方法
    this.props.handleDelete(this.props.index);
  }
}

ToDoItem.propTypes = {  //限制父組件傳入值的型別，不強制限制傳錯會報錯
  item: propTypes.oneOfType([propTypes.string,propTypes.number]), //輸入的值是其中一種型別
  handleDelete: propTypes.func,
  index: propTypes.number
}

ToDoItem.defaultProps = {  //定義如果沒有傳遞值進來的話給的默認值
  item: 'HI'
}
export default ToDoItem;
