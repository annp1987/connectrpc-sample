import { useState } from 'react'
import Linkify from "linkify-react";
import './App.css'

import { createPromiseClient } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";

// Import service definition that you want to connect to.
import { GreetService } from './proto/greet_connect.js';

// The transport defines what type of endpoint we're hitting.
// In our example we'll be communicating with a Connect endpoint.
const transport = createConnectTransport({
  baseUrl: "http://localhost:8080",
});

// Here we make the client itself, combining the service
// definition with the transport.
const client = createPromiseClient(GreetService, transport);

function App() {
  const [phoneNumberInput, setPhoneNumber] = useState('');
  const [messageInput, setMessage] = useState('');

  const [msgResponse, setRespMessages] = useState<{
      prefix: string;
      message: string;
  }>();
  return <>
    <div>
      <p>{msgResponse?.prefix}</p>
      <Linkify>{msgResponse?.message}</Linkify>
    </div>
    <form onSubmit={async (e) => {
      e.preventDefault();
      // Clear inputValue since the user has submitted.
      setPhoneNumber("");
      setMessage("");

      const response = await client.greet({
        number: phoneNumberInput,
        message: messageInput,
      });
      setRespMessages(
        {
          prefix: JSON.stringify(response.prefix),
          message: response.message,
        });
    }}>
      <div>
        <label htmlFor="phone-number">Phone Number:</label><br/>
        <input
          id="phone-number"
          value={phoneNumberInput}
          onChange={(event) => setPhoneNumber(event.target.value)}
        /><br/>
        <label htmlFor="message">Message:</label><br/>
        <textarea
          id="message"
          value={messageInput}
          onChange={(event) => setMessage(event.target.value)}
        /><br/>
        <button type="submit">Send</button>
      </div>
    </form>
    <div>

    </div>
  </>;
}

export default App
