import React from 'react'
import { Route, Switch } from 'react-router'

import Main from './Main'
import About from './About'
import { BrowserRouter } from 'react-router-dom'

function App() {
  return (
    <BrowserRouter>
      <div>
        <Switch>
          <Route path="/" exact component={Main} />
          <Route exact path="/about" component={About} />
        </Switch>
      </div>
    </BrowserRouter>
  )
}

export default App
