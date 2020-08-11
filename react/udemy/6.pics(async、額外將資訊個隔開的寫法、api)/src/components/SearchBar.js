import React from 'react';

class SearchBar extends React.Component{
    constructor(props) {
        super(props);
        this.handleChange = this.handleChange.bind(this);
        this.onForSubmit = this.onForSubmit.bind(this);
      }

    onForSubmit=(event)=>{
        event.preventDefault(); //來避免瀏覽器預設行為，如HTML的return false
        this.props.onSearchSubmit(this.state.term);//取外面傳入的函數
    }

    handleChange(e){
        this.setState({term : e.target.value});
    }

    state={term:''};

    render(){
        return (
            <div className='ui segment'>
                <form onSubmit={this.onForSubmit} className='ui form'>
                    <div className='field'>
                        <label>Image search</label>
                        <input type="text" value={this.state.term} onChange={this.handleChange}/>
                        <button type="submit">submit</button>
                    </div>
                </form>
            </div>
        );
    }
}


export default SearchBar;