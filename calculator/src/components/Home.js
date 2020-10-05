import React, { Component } from "react";
import {Container, Button, Form, FormGroup, Input, Table} from "reactstrap";

import axios from "axios";

import { API_URL } from "../constants";

class Home extends Component {
  state = {
    text: "0, 0, 0, 0, 0, 0, 0, 0, 0, 0",
    arr: [0,0,0,0,0,0,0,0],
    data: []
  };

  onChange = e => {
    this.setState({[e.target.name]: e.target.value});
  }

  sendArr = e => {
    e.preventDefault();
    const arrs = this.state.text.split(",").map(x=>+x);
    this.state.arr = arrs.filter(value => !Number.isNaN(value));
    // console.log(this.state.arr)
    axios.put(API_URL + "/1", this.state)
      .then(res => {
        const arr1 = res.data.arr;
        const arr2 = res.data.data;
        if (arr2[arr2.length-1] == null) {
          arr2[arr2.length - 1] = 'Each element appears exactly once'
        } else {
          arr2[7] = arr2[7] + " (appears "+ arr2[arr2.length - 1] + " times)"
        }
        console.log(arr2)
        this.setState({
          data: arr2,
          arr: arr1,
          text: arr1.toString()
        })
      });
  }


  render() {
    return (
      <Container style={{ marginTop: "20px" }} maxwidth="sm">
        <div align="center">
          <Form onSubmit={this.sendArr}>
            <p align="left">Please enter a list of numbers seperated by commas below and press calculate to get the results:</p>
            <FormGroup>
              <Input
                type = "textarea"
                name = "text"
                onChange={this.onChange}
                style={{height: "150px"}}
                placeholder={this.state.text}
              >
              </Input>
            </FormGroup>
            <Button 
              className="float center"
              color="danger"
              style={{ minWidth: "200px" }, {height: "50px"}}
              >
                Calculate
            </Button>
          </Form>
          <hr />
            <h1 style={{marginTop: "20px"}}> Results </h1>
            <Table dark bordered striped>
              <thead>
                <tr>
                  <th>Category</th>
                  <th>Data</th>
                </tr>
              </thead>
              <tbody>
                <tr>
                  <td>Mean</td>
                  <td>{this.state.data[0]}</td>
                </tr>
                <tr>
                  <td>Standard Deviation</td>
                  <td>{this.state.data[1]}</td>
                </tr>
                <tr>
                  <td>Minimum</td>
                  <td>{this.state.data[2]}</td>
                </tr>
                <tr>
                  <td>25th Percentile</td>
                  <td>{this.state.data[3]}</td>
                </tr>
                <tr>
                  <td>Median (50th Percentile)</td>
                  <td>{this.state.data[4]}</td>
                </tr>
                <tr>
                  <td>75th Percentile</td>
                  <td>{this.state.data[5]}</td>
                </tr>
                <tr>
                  <td>Maximum</td>
                  <td>{this.state.data[6]}</td>
                </tr>
                <tr>
                  <td>Mode</td>
                  <td>{this.state.data[7]}</td>
                </tr>
                <tr>
                  <td>Sorted Array</td>
                  <td>{this.state.text}</td>
                </tr>
              </tbody>
            </Table>
        </div>
      </Container>
    );
  }
}

export default Home;