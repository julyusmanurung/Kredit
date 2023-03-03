import Button from 'react-bootstrap/Button';
import Modal from 'react-bootstrap/Modal';
import ProfilePicture from '../../../assets/profile.jpg'
import {Col, Image, Row} from "react-bootstrap";
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import {faAddressBook, faUser, faEnvelope, faChair} from '@fortawesome/free-solid-svg-icons'
import React, {useEffect, useState} from "react";
import axios from "axios";

const UserProfile = (props) => {
    const localhost = "http://localhost:8080/"
    const [userProfile, setUserProfile] = useState(null)

    const getUserProfile = async () => {
        const res = await axios.get(localhost + "profile/" + localStorage.getItem('nik'))
        setUserProfile(res.data)
    }

    useEffect(() => {
        getUserProfile()
    }, [])

    return (
        <Modal
            {...props}
            size="lg"
            aria-labelledby="contained-modal-title-vcenter"
            centered
        >
            <Modal.Header closeButton>
                <Modal.Title id="contained-modal-title-vcenter">
                    User Profile
                </Modal.Title>
            </Modal.Header>
            <Modal.Body className="d-flex gap-4 align-items-center modal-background">
                <Image src={ProfilePicture} className="modal-profile-picture"></Image>
                <Row className="gap-4">
                    <Row>
                        <Col className="d-flex flex-column">
                            <label className="description"><FontAwesomeIcon icon={faAddressBook} className="pe-2"/>NIK</label>
                            <label>{!userProfile ? "Loading": userProfile.data.user_id}</label>
                        </Col>
                    </Row>
                    <Row>
                        <Col className="d-flex flex-column">
                            <label className="description"><FontAwesomeIcon icon={faUser} className="pe-2"/>Nama</label>
                            <label>{!userProfile ? "Loading": userProfile.data.name}</label>
                        </Col>
                    </Row>
                    <Row>
                        <Col className="d-flex flex-column">
                            <label className="description"><FontAwesomeIcon icon={faEnvelope} className="pe-2"/>Email</label>
                            <label>{!userProfile ? "Loading": userProfile.data.email}</label>
                        </Col>
                    </Row>
                    <Row>
                        <Col className="d-flex flex-column">
                            <label className="description"><FontAwesomeIcon icon={faChair} className="pe-2"/>Level - Jabatan</label>
                            <label>{!userProfile ? "Loading": userProfile.data.level} - {!userProfile ? "Loading": userProfile.data.jabatan}</label>
                        </Col>
                    </Row>
                </Row>
            </Modal.Body>
            <div className="text-center mb-3">
                <Button variant="sinarmas2" onClick={props.onHide} className="w-25">Close</Button>
            </div>
            {/*<Modal.Footer className="d-flex justify-content-center">*/}
            {/*</Modal.Footer>*/}
        </Modal>
    );
}

export default UserProfile