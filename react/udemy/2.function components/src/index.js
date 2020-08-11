// Import the React and ReactDOM libraries
import React from 'react';
import ReactDOM from 'react-dom';
import faker from 'faker';
import CommentDetail from './CommentDetail';//import ./CommentDetail的公開函式CommentDetail
//Create a react component
import ApprovalCard from './ApprovalCard';

const App = function(){
    return <div className="ui container comments">
        <ApprovalCard>
            <div>
                <h4>warning</h4>
                Are you sure you want to
            </div>
        </ApprovalCard>


        <ApprovalCard>
            <CommentDetail date="Today at 7:00PM" content={faker.lorem.text()}/>
        </ApprovalCard> 
        <ApprovalCard>
            <CommentDetail date="Today at 8:00PM" content={faker.lorem.text()}/>
        </ApprovalCard> 
        <ApprovalCard>
            <CommentDetail date="Today at 9:00PM" content={faker.lorem.text()}/>
        </ApprovalCard> 
       
    </div>
};

//react的style必須用{{}}包住，使用函式裡的變數要用{}包住



//Take the react component and show it on the screen
ReactDOM.render(<App />, document.querySelector("#root"));//ReactDOM.render(element, container[, callback])



