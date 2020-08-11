// Import the React and ReactDOM libraries
import React from 'react';
import ReactDOM from 'react-dom';
import TEST from './App';

function gettime(){
    return (new Date()).toLocaleTimeString();
}

//Create a react component
const App = function(){
    return <div>hi there!</div>
};

const App2 = ()=>{
    const buttonText = {text:'Click me!'};
    const style = {background:'blue',color:'white'}

    return (
    <div>
        <App/>
        <TEST/>
        <label className="label" htmlFor="name">Enter name:</label>
        <input type="text" id="name"/>

        <button style={{background:'blue',color:'white'}}>
        {gettime()}
        </button> 

        <button style={style}>
        {buttonText.text}
        </button> 
    
    </div>
    );
};
//react的style必須用{{}}包住，使用函式裡的變數要用{}包住



//Take the react component and show it on the screen
ReactDOM.render(<App2 />, document.querySelector("#root"));



