import React from 'react';

class SearchBar extends React.Component{
    constructor(props) {
        super(props);
        this.onInputChange = this.onInputChange.bind(this);
        this.onFormSubmit = this.onFormSubmit.bind(this);
      }

    state = {term:''};

    onInputChange(event){
        this.setState({term:event.target.value});
    };

    onFormSubmit(event){ 
        event.preventDefault(); //等於return false
        this.props.onFormSubmit(this.state.term);

    };

    render(){
        return (
        <div className="search-bar ui segment">
           <form onSubmit={this.onFormSubmit} className="ui form">
               <div className="field">
                   <label>video search</label>
                   <input 
                   type="text" 
                   value={this.state.term}
                   onChange={this.onInputChange}
                   />
               </div>
               <button type="submit">submit</button>
           </form>
        </div>
    )}
}

export default SearchBar;