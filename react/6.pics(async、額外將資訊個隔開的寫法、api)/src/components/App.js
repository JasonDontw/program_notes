import React from 'react';
import axios from 'axios';
import SearchBar from './SearchBar';
import unsplash from '../api/unsplash';
import ImageList from './ImageList';

class App extends React.Component {
    
    constructor(props) {
        super(props);
        this.onSearchSubmit = this.onSearchSubmit.bind(this);
      }

    state = {images : []};
    //正常寫法
   /*
    onSearchSubmit(term){
        axios.get('https://api.unsplash.com/search/photos',{
            params:{query : term},
            headers: {
                Authorization:
                'Client-ID 6c1075170bebd4b6eb7f9f939482f24c90c87704c38a8c85604fd1a53d6e1c63',
            }
        }).then((response) =>{
            this.setState({image:response.data.results});
        });

    }
    */

    //async 寫法(語法糖) 
    /*async onSearchSubmit(term){
        const response = await axios.get('https://api.unsplash.com/search/photos',{
            params:{query : term},
            headers: {
                Authorization:
                'Client-ID 6c1075170bebd4b6eb7f9f939482f24c90c87704c38a8c85604fd1a53d6e1c63',
            }
        });
        
        this.setState({images:response.data.results});
        console.log(this.state.images)
    }
    */

    //額外將資訊個隔開的寫法
   async onSearchSubmit(term){
    const response = await unsplash.get('/search/photos',{
        params:{query : term},
       
    });
    
    this.setState({images : response.data.results});
    
}
    

    render(){
        return (
            <div className='ui container' style={{marginTop:'10px'}}>
                <SearchBar onSearchSubmit={this.onSearchSubmit}/>
                <ImageList images={this.state.images}/>
            </div>
            );
    }
};

export default App;