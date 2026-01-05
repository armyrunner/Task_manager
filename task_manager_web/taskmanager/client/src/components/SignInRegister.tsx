import Button  from 'react-bootstrap/Button';
import Form  from 'react-bootstrap/Form';

function SignIn() {
    return (
        <div className='content justify-content-center alignitems-center d-flex shadow-lg' id='content'>
            <div className='col-md-6 d-flex justify-content-center'>
                <Form>
                    <div className="header-text mb-4">
                        <h1>Create Account</h1>
                    </div>
                    <div className="input-group mb-3">
                        <input type="text" placeholder='Name' className='form-control form-control-lg bg-light fs-6'></input>
                    </div>
                    <div className="input-group mb-3">
                        <input type="email" placeholder='Email' className='form-control form-control-lg bg-light fs-6'></input>
                    </div>
                    <div className="input-group mb-3 justify-content-center" >
                        <input type="password" placeholder='Password' className='form-control form-control-lg bg-light fs-6'></input>
                    </div>
                </Form>
            </div>
        </div>
     );
}

export default SignIn;