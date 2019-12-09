export const selectSong = (song) => { //type: 每個action都會有type屬性來描述state該怎麼改變。    
    return{                           
        type:'SONG_SELECTED', 
        payload:song
    };
};
