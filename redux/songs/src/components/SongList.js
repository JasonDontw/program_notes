import React from 'react';
import {connect} from 'react-redux';
import {selectSong} from '../actions' 
class SongList extends React.Component{
    renderList(){
        return this.props.songs.map((song)=>{//因connect所以可以取用mapStateToProps所return的songs
            return(
                <div className='item' key={song.title}>
                    <div className='right floated content'>
                        <button className='ui button primary' onClick={()=>this.props.selectSong(song)}>
                            select
                        </button>
                    </div>

                    <div className='content'>{song.title}</div>
                </div>
            );
        });
    }

    render(){
        return <div className='ui divided list'>{this.renderList()}</div>
    }
};

const mapStateToProps = (state) =>{ //名稱可自取，因用react-redux的connect綁定的關係會直接將redux裡reducers的參數全數傳入
    console.log(state);
    return { songs : state.songs };
};


export default connect(
    mapStateToProps, //將mapStateToProps綁定到SongList，，讓renderList能取用songs
    { selectSong } //將action綁定到SongList，，讓renderList能取用selectSong
    )(SongList);