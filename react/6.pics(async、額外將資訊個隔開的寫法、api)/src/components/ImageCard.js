import React from 'react';

class ImageCard extends React.Component{
   constructor(props){
       super(props);
       this.state = { spans:0 }
       this.imageRef = React.createRef();
   }
   componentDidMount(){
       this.imageRef.current.addEventListener('load',this.setSpans);
    //因為直接呼叫this.imageRef.current.clientHeight是全部的current
    //所以不會存clientHeight，所以要在他新增的時候就用監聽一個一個先取出
       console.log(this.imageRef.current);
   }

   setSpans = () => {
      const height = this.imageRef.current.clientHeight;
      const spans = Math.ceil(height / 10 + 1);
      this.setState({spans: spans});
   }

    render(){
        return (
            <div style={{gridRowEnd:`span ${this.state.spans}`}}>
                <img 
                alt={this.props.image.description}
                src={this.props.image.urls.regular}
                ref={this.imageRef}
                />
            </div>
        );
    }
}

export default ImageCard;