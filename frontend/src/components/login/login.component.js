import React from 'react'
import {Button, Container, Form, Row, InputGroup, Image} from 'react-bootstrap'
import './login.component.css'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import {faIdCardClip, faLock} from '@fortawesome/free-solid-svg-icons'
import Logo from '../../assets/logo.png'
import Swal from "sweetalert2"
import axios from "axios";

class LoginForm extends React.Component {
    constructor(props) {
    super(props)
    this.state = {
      validated: false
    }
  }
  render() {
    return (
        <Container fluid className="login-component">
          <Form noValidate validated={this.state.validated} onSubmit={this.handleSubmit}>
            <Form.Group className="form-header">
              <Image src={Logo} className="logo-bank-sinarmas"></Image>
            </Form.Group>
            <Row className="mb-3">
              <Form.Group>
                <Form.Label>NIK <sup>*</sup></Form.Label>
                <InputGroup hasValidation>
                  <InputGroup.Text id="inputGroupPrepend"><FontAwesomeIcon icon={faIdCardClip}/></InputGroup.Text>
                  <Form.Control
                      name="userid"
                      type="text"
                      placeholder="NIK"
                      aria-describedby="inputGroupPrepend"
                      required
                  />
                  <Form.Control.Feedback type="invalid">
                    User ID fields is required
                  </Form.Control.Feedback>
                </InputGroup>
              </Form.Group>
            </Row>
            <Row className="mb-3">
              <Form.Group>
                <Form.Label>Password <sup>*</sup></Form.Label>
                <InputGroup hasValidation>
                  <InputGroup.Text id="inputGroupPrepend"><FontAwesomeIcon icon={faLock}/></InputGroup.Text>
                  <Form.Control
                      name="password"
                      type="password"
                      placeholder="Password"
                      aria-describedby="inputGroupPrepend"
                      required
                  />
                  <Form.Control.Feedback type="invalid">
                    Password field is required.
                  </Form.Control.Feedback>
                </InputGroup>
              </Form.Group>
            </Row>
            <Button type="submit" className="w-100" variant="sinarmas">Login</Button>
          </Form>
        </Container>
  )
  }

  handleSubmit = async (event) => {
    event.preventDefault()
    const formData = new FormData(event.currentTarget)
    const body = {
      userid: formData.get("userid"),
      password: formData.get("password")
    }

    const localhost = "http://localhost:8080/"
    const res = await axios.get(localhost + "login?nik=" + body.userid + "&password=" + body.password)

    if (res.data.message === "login succcess") {
      Swal.fire({
        icon: 'success',
        title: 'You are logged in',
        showConfirmButton: false,
        timer: 1500
      }).then( () => {
          localStorage.setItem('user', "true")
          localStorage.setItem('sidebar', 'open')
          localStorage.setItem('mode', 'light')
          localStorage.setItem('nik', res.data.data.nik)
          window.location.href="/"
        }
      )
    } else {
      Swal.fire({
        icon: 'error',
        title: 'Oops...',
        text: 'Invalid nik or password',
      })
    }

    const form = event.currentTarget
    if (form.checkValidity() === false) {
      event.preventDefault()
      event.stopPropagation()
    }

    this.setState({
      validated: true
    })

  }

}

export default LoginForm
