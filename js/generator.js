'use strict';

const creator = React.createElement;


class MainContent extends React.Component {
	constructor(props){
		super(props);
	}

	render(){
		return creator(
			'div',
			{onClick: () => { console.log("hello")}},
			"content to implement"
		)
	}
}


const domContainer = document.querySelector('#root');
ReactDOM.render(creator(MainContent), domContainer);