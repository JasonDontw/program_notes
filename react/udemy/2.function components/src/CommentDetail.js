import React from 'react';
import faker from 'faker';

const CommentDetail = (props) =>{ //props為call CommentDetail時傳入的變數
    return (
            <div className="comment">
                <a href="/" className="avatar">
                <img alt="avatar" src={faker.image.avatar()}/>
                </a>
                <div className="content">
                <a href="/" className="author">
                {faker.internet.userName()}
                </a>
                <div className="metadata">
                    <span className="Date">{props.date}</span>
                </div>
                <div className="text">{props.content}</div>
                </div>
                <hr/>
            </div>
    );
};

export default CommentDetail;//宣告要公開的函式