import './VideoItem.css'
import React from 'react';


const VideoItem = props =>{
    return (
         //此處onClick必須用匿名函式呼叫函式不能直接使用
        <div onClick={()=>props.onVideoSelect(props.video)} className="video-item item">
            <img alt={props.video.snippet.title} className="ui image" src={props.video.snippet.thumbnails.medium.url}/>
            <div className="content">
                <div className="header">{props.video.snippet.title}</div>
            </div>
        </div>
    )
};

export default VideoItem;