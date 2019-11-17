import React from 'react';
import ReactDOM from 'react-dom';


// function success(pos) {
//   console.log(pos)
// };


class App extends React.Component{
constructor(props){
  super(props);
  /*super 會參照父類別的建構子。
    (在我們的例子當中，它會指向 React.Component 的實作。)
    直到你呼叫父類別的建構子後，你才能在建構子中使用 this
    若是需要綁定 this.方法或是需要在 constructor 使用 props，定義 state，就需要 constructor。
    若是在其他方法（如 render）使用 this.props 則不用一定要定義 constructor*/
  this.state = {lat: null, long: null, errorMessage:''}; //內建的功能

  window.navigator.geolocation.getCurrentPosition(
    (position) => {
      this.setState({lat: position.coords.latitude})
    },
    (err) => {
      this.setState({errorMessage:err.message})
    }
  );
//寫法2 先定義SUCCESS函式，再直接使用
    /*window.navigator.geolocation.getCurrentPosition(success);*/
}

  render(){
    if(this.state.errorMessage && !this.state.lat){
      return (
        <div>
             err:{this.state.errorMessage}
        </div>
        
      );
    }
    if(!this.state.errorMessage && this.state.lat){
      return (
        <div>
            Latitude:{this.state.lat}
        </div>
        
      );
    }
    if(!this.state.errorMessage && !this.state.lat){
      return (
        <div>
             Loading!
        </div>
        
      );
    }
   
    
  }

}

ReactDOM.render(<App />,document.querySelector('#root'));