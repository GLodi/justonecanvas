import React from 'react'
import './About.css'

function About() {
  return (
    <body className="about">
      <div className="center">
        <h1>
          <p> literally just one canvas</p>
        </h1>
        <ul className="nobull">
          <li>
            ...but a bit special.
            <br />
            &nbsp;
          </li>
          <li>
            This is essentially a copy of Reddit's 2017 r/place. Just tinier
            (65x65) and designed to work on a 5$ Digital Ocean droplet.
          </li>
          <li>
            All users connected to justonecanvas actually share the same canvas,
            and the backend
          </li>
          <li>
            updates all players of each other's moves thanks to WebSockets.
            <br />
            &nbsp;
          </li>

          <li>
            You can move once every second.
            <br />
            &nbsp;
          </li>

          <li>
            <a href="https://giuliolodi.dev/2020/08/29/how-i-made-justonecanvas">Blog Post</a>
          </li>
        </ul>
      </div>
    </body>
  )
}

export default About
