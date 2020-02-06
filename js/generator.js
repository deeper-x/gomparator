'use strict';

const creator = React.createElement;

class MainContent extends React.Component {
	constructor(props) {
		super(props);
		this.state = {value: ""}

		this.handleChange = this.handleChange.bind(this)
		this.handleSubmit = this.handleSubmit.bind(this)
	}

	urlIsValid(inputUrl) {
		// TODO
		return true;
	}

	createUrl (inputUrl) {
		// TODO - check url is valid
		const elements = inputUrl.split("/");
		return `https://api.github.com/repos/${elements[3]}/${elements[4]}`; 
	}

	storeRecord(stringToSave){
		// axios.get(
		// 	`http://127.0.0.1:8000/?query={saveData(data:${stringToSave})}`
		// ).then(function(){
		// 	console.log(`TODO saving on redis ${stringToSave}`);
		// })
	}

	handleChange(e) {
		this.setState({value: e.target.value});
	}

	handleSubmit(e) {
		const url = this.createUrl(this.state.value);
		
		axios.get(url)
		.then((response)=> {
			ReactDOM.render(
				<div>Name: { response.data.name} - Stars { response.data.stargazers_count}</div>,
				document.getElementById("result")
			);
			this.storeRecord(response.data.name);
		});
		
		e.preventDefault();
	}

	render() {
		return (
			<div>
				<form onSubmit={this.handleSubmit}>
					<input type="text" value={this.state.value} id="github_uri" onChange={this.handleChange} />
					<input type="submit" />
				</form>
			</div>
		)
	}
}


const domContainer = document.querySelector('#root');
ReactDOM.render(creator(MainContent), domContainer);