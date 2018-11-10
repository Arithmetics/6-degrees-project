import React, { Component } from 'react';
import './App.css';
import './TermsList.js'
import TermsList from './TermsList.js';

const API = 'http://192.168.0.5:8080/api/fibonacci';
const EXAMPLENUM = 7;


class App extends Component {
  constructor(props){
    super(props);

    this.state = {
      result: null,
      loading: false,
      terms: 7,
    };


  }

  componentDidMount(){
    
    this.calcData(EXAMPLENUM);
    
  }

  calcData = (num1) => {
    this.setState({ loading: true });
    fetch(`${API}/${num1}`)
    .then(response => response.json())
    // .then(resp => console.log(resp))
    .then(resp => this.setState({result: resp, loading: false}));
  }

  handleChange = (e) => {
    this.setState({terms: e.target.value});
  }

  handleSubmit = (e) => {
    e.preventDefault();
    this.calcData(this.state.terms, this.state.numTwo);
  }

  render() {
    const result = this.state.result;
    const loading = this.state.loading;

    
      return (
        <div className="App">
        <div className="top">
          <h1>Fibonacci Sequence Generator</h1>
          <h3>1. Enter the number of terms you would like to generate</h3>
          <h3>2. Click Generate</h3>
          <h3>3. Enjoy!</h3>
        </div>
          <form onSubmit={this.handleSubmit}>
            <label>
              Terms:
            </label>
            <input type="number" min="0" max="95" value={this.state.terms} onChange={this.handleChange} />
            <input className="submit" type="submit" value="Generate" ></input>
          </form>
          <TermsList results={this.state.result}/>
        </div>
      );
    
  }
}

export default App;
