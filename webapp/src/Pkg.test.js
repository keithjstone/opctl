import React from 'react';
import ReactDOM from 'react-dom';
import Pkg from './Pkg';

it('renders without crashing', () => {
    /* arrange */
    const div = document.createElement('div');

    /* act/assert */
    ReactDOM.render(<Pkg value={{}}/>, div);
});
