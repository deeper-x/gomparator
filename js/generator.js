'use strict';

const creator = React.createElement;


class MainContent extends React.Component {
	constructor(props){
		super(props);
	}

	render(){
		return (
			<div>Hello from JSX!</div>
		)
	}
}


const domContainer = document.querySelector('#root');
ReactDOM.render(creator(MainContent), domContainer);