import Button from 'react-bootstrap/Button';
import Modal from 'react-bootstrap/Modal';
import {Col, Form, FormControl, FormGroup, FormLabel, Row} from "react-bootstrap"
import axios from "axios"
import Swal from "sweetalert2"
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import {faCircleXmark, faUserCircle, faCircleCheck} from '@fortawesome/free-solid-svg-icons'
import {useState} from "react";

const ChangePassword = (props) => {
    const localhost = "http://localhost:8080/"

    const requiredPassed = <FontAwesomeIcon icon={faCircleCheck}/>
    const requredNotPassed = <FontAwesomeIcon icon={faCircleXmark}/>
    const [submitted, setSubmitted] = useState(false);

    const [charLengthCase, setCharLengthCase] = useState(false)
    const [lowercase, setLowercase] = useState(false)
    const [uppercase, setUppercase] = useState(false)
    const [numeric, setNumeric] = useState(false)
    const [special, setSpecial] = useState(false)

    const changePasswordValidation = (str) => {
        const lowerCaseRegex = /[a-z]/;
        const upperCaseRegex = /[A-Z]/;
        const numericRegex = /[0-9]/;
        const specialCharRegex = /[@#$%^&+=]/;

        if (str.length < 8) {
            setCharLengthCase(false)
            return false
        } else {
            setCharLengthCase(true)
        }

        if (!lowerCaseRegex.test(str)) {
            setLowercase(false)
            return false
        } else {
            setLowercase(true)
        }

        if (!upperCaseRegex.test(str)) {
            setUppercase(false)
            return false
        } else {
            setUppercase(true)
        }

        if (!numericRegex.test(str)) {
            setNumeric(false)
            return false
        } else {
            setNumeric(true)
        }

        if (!specialCharRegex.test(str)) {
            setSpecial(false)
            return false
        } else {
            setSpecial(true)
        }

        return true
    }

    const checkConfirmPassword = (newPassword, confirmPassword) => {
        return newPassword === confirmPassword
    }

    const handleSubmit = async (event) => {
        event.preventDefault()

        setSubmitted(true)

        const formData = new FormData(event.currentTarget)
        const body = {
            oldPassword: formData.get("oldpassword"),
            newPassword: formData.get("newpassword"),
            confirmPassword: formData.get("confirmpassword")
        }
        const isNewPasswordValid = changePasswordValidation(body.newPassword)

        if (!isNewPasswordValid) {
            console.log("Your new password does not meet the requirements")
             Swal.fire({
                icon: 'error',
                title: 'Oops...',
                text: 'Your new password does not meet the requirements',
            })
        } else {
            if (checkConfirmPassword(body.newPassword, body.confirmPassword)) {
                const res = await axios.patch(localhost + "updatepassword/" + localStorage.getItem("nik"), {old_password: body.oldPassword, new_password: body.newPassword})
                if (res.data.message === "password berhasil diubah") {
                     Swal.fire({
                        icon: 'success',
                        title: 'Your password has been changed successfully',
                        text: 'You have to log in again',
                    }).then( () => {
                        localStorage.removeItem('user')
                        localStorage.removeItem('sidebar')
                        localStorage.removeItem('mode')
                        localStorage.removeItem('nik')
                        window.location.href="/"
                    })
                } else {
                    console.log(res.data.message)
                    Swal.fire({
                        icon: 'error',
                        title: 'Oops...',
                        text: 'Your old password is wrong',
                    })
                }
            } else {
                Swal.fire({
                    icon: 'error',
                    title: 'Oops...',
                    text: 'Your confirm password does not match with the new password',
                })
            }
        }
    }

    return (
        <Modal
            {...props}
            size="lg"
            aria-labelledby="contained-modal-title-vcenter"
            centered
            className="change-password-modal"
        >
            <Modal.Header closeButton>
                <Modal.Title id="contained-modal-title-vcenter">
                    Change Password
                </Modal.Title>
            </Modal.Header>
            <Modal.Body>
                <Form onSubmit={handleSubmit}>
                    <Row className="align-items-center px-2">
                        <Col className="d-flex flex-column gap-3">
                            <FormGroup>
                                <FormLabel className="required">Current Password</FormLabel>
                                <FormControl required name="oldpassword" type="password" placeholder="Your latest password"></FormControl>
                            </FormGroup>
                            <FormGroup>
                                <FormLabel className="required">New Password</FormLabel>
                                <FormControl required name="newpassword" type="password" placeholder="Your new password"></FormControl>
                            </FormGroup>
                            <FormGroup>
                                <FormLabel className="required">Confirm Password</FormLabel>
                                <FormControl required name="confirmpassword" type="password" placeholder="Retype your new password"></FormControl>
                            </FormGroup>
                        </Col>
                        <Col>
                            <h4 className="mb-3">Password requires: </h4>
                            {submitted ?
                                <ul className="pass-req p-0">
                                    <li className={charLengthCase ? "true" : "false"}>{charLengthCase ? requiredPassed: requredNotPassed} At least 8 characters in length</li>
                                    <li className={lowercase ? "true" : "false"}>{lowercase ? requiredPassed: requredNotPassed} 1 lower case letter [a-z]</li>
                                    <li className={uppercase ? "true" : "false"}>{uppercase ? requiredPassed: requredNotPassed} 1 upper case letter [A-Z]</li>
                                    <li className={numeric ? "true" : "false"}>{numeric ? requiredPassed: requredNotPassed} 1 numeric character [0-9]</li>
                                    <li className={special ? "true" : "false"}>{special ? requiredPassed: requredNotPassed} 1 special character</li>
                                </ul> :
                                <ul className="pass-req p-0">
                                    <li>At least 8 characters in length</li>
                                    <li>1 lower case letter [a-z]</li>
                                    <li>1 upper case letter [A-Z]</li>
                                    <li>1 numeric character [0-9]</li>
                                    <li>1 special character</li>
                                </ul>
                            }
                        </Col>
                    </Row>
                    <div className="d-flex justify-content-center">
                        <Row className="mb-2 mt-4 justify-content-center gap-3" style={{width: "75%"}}>
                            <Button type="submit" variant="sinarmas" className="w-50">Ganti Password</Button>
                            <Button variant="sinarmas2" onClick={props.onHide} className="w-25">Close</Button>
                        </Row>
                    </div>
                </Form>
            </Modal.Body>
            {/*<Modal.Footer className="d-flex justify-content-center">*/}
            {/*</Modal.Footer>*/}
        </Modal>
    );
}

export default ChangePassword