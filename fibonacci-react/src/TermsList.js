import React, { Component } from 'react';
import './App.css';
import './TermsList.js'

const API = 'http://192.168.0.5:8080/api/fibonacci';
const EXAMPLENUM = 7;


class TermsList extends Component {
  constructor(props){
    super(props);
  }

  

  render() {
    const results = this.props.results;
    const resultString = JSON.stringify(results);
    const final = resultString.substring(1,resultString.length-1)
    return (
      <div className="terms">
        <p>{final}</p>
      </div>
    )
  }
}

export default TermsList;
