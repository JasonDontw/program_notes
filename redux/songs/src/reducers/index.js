//Reducer其實只是javascript的function，而這個function內含兩個參數，
//第一個為current state，也就是目前state的資料，另一個則是action，
//如果我們想要更改store內的state資料，則要透過action的幫忙。
import {combineReducers} from 'redux';


const songsReducer = () => {
    return[
        {title:'no Scrubs',duration:'4:05'},
        {title:'Macarena',duration:'2:35'},
        {title:'All Star',duration:'3:15'},
        {title:'I Want it That Way',duration:'1:45'}
    ];
};


const selectedSongReducer = (selectedSong=null , action) => { //設定reducer function，第一個參數為state，第二個參數為action任務
    if(action.type === 'SONG_SELECTED'){
        return action.payload;
    }

    return selectedSong;
};


export default combineReducers({
    songs: songsReducer,
    selectedSong:selectedSongReducer
});
//CombineReducers輔助函數的作用是，把一個由多個不同的reducer函數作為值的對象，合併成一個最終的reducer函數，然後就可以對這個reducer調用createStore。 


