import React, {Component} from 'react';

export default class SocketInput extends Component {
  constructor(props) {
    super(props);

    this.state = {
      value: null,
    };

    this.handleArgChange = this.handleArgChange.bind(this);
  }

  handleArgChange(e) {
    const value = e.target.value;
    this.props.onArgChange({socket: value});
    this.setState({value});
  };

  render() {
    return (
      <div className='form-group'>
        <label className='form-control-label' htmlFor={this.props.name}>{this.props.name}</label>
        <p className='custom-control-description'>{this.props.socket.description}</p>
        <input
          className='form-control'
          id={this.props.name}
          onChange={this.handleArgChange}
          placeholder='/absolute/path/of/socket'
          type='text'
          value={this.state.value}
        />
      </div>
    );
  }
}
