// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

// https://www.iacr.org/cryptodb/archive/2002/ASIACRYPT/50/50.pdf
contract Verification
{
    uint256 constant CURVE_ORDER = 0x30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001;
    uint256 constant CURVE_B = 3;

    // a = (p+1) / 4
    uint256 constant CURVE_A = 0xc19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f52;

    struct G1Point {
        uint X;
        uint Y;
    }

    // Encoding of field elements is: X[0] * z + X[1]
    struct G2Point {
        uint[2] X;
        uint[2] Y;
    }

    // (P+1) / 4
    function A() pure internal returns (uint256) {
        return CURVE_A;
    }
/*
    function P() pure internal returns (uint256) {
        return FIELD_ORDER;
    }
*/

    function N() pure internal returns (uint256) {
        return CURVE_ORDER;
    }

    /// return the generator of G1
    function P1() pure internal returns (G1Point memory) {
        return G1Point(1, 2);
    }

    /// return the sum of two points of G1
    function g1add(G1Point memory p1, G1Point memory p2) view internal returns (G1Point memory r) {
        uint[4] memory input;
        input[0] = p1.X;
        input[1] = p1.Y;
        input[2] = p2.X;
        input[3] = p2.Y;
        bool success;
        assembly {
            success := staticcall(sub(gas(), 2000), 6, input, 0xc0, r, 0x60)
        // Use "invalid" to make gas estimation work
        //switch success case 0 { invalid }
        }
        require(success);
    }

    /// return the product of a point on G1 and a scalar, i.e.
    /// p == p.mul(1) and p.add(p) == p.mul(2) for all points p.
    function g1mul(G1Point memory p, uint s) view internal returns (G1Point memory r) {
        uint[3] memory input;
        input[0] = p.X;
        input[1] = p.Y;
        input[2] = s;
        bool success;
        assembly {
            success := staticcall(sub(gas(), 2000), 7, input, 0x80, r, 0x60)
        // Use "invalid" to make gas estimation work
        //switch success case 0 { invalid }
        }
        require(success);
    }

/*
    function VSSVerify(G1Point[] memory commitments) public payable returns (bool)
    {
        for(uint i=0;i<Gs.length;i++){
            G1Point memory right;
            uint256 xPower=1;
            uint256 x=i+1;
            G1Point memory temp;
            for(uint j=0;j<commitments.length;j++){
                temp = g1mul(commitments[j], xPower);
                right=g1add(right, temp);
                xPower=mulmod(xPower, x, CURVE_ORDER);
            }
            if(Gs[i].X != right.X||Gs[i].Y != right.Y){
                VerificationResult.push(false);
                return false;
            }
        }
        VerificationResult.push(true);
        return true;
    }

    function DLEQVerifyCKeys() public payable returns (bool)
    {
        for (uint256 i = 0; i < DLEQProofCKeys.length; i++)
        {
            G1Point memory gG = g1mul(g, DLEQProofCKeys[i].z);
            G1Point memory y1G = g1mul(Gs[i], DLEQProofCKeys[i].c);

            G1Point memory hG = g1mul(g1add(ciphertext.C0, TTPsPk[i]), DLEQProofCKeys[i].z);
            G1Point memory y2G = g1mul(CKeys[i], DLEQProofCKeys[i].c);

            G1Point memory pt1 =  g1add(gG, y1G);
            G1Point memory pt2 =  g1add(hG, y2G);
            if ((DLEQProofCKeys[i].a1.X != pt1.X) || (DLEQProofCKeys[i].a1.Y != pt1.Y) || (DLEQProofCKeys[i].a2.X != pt2.X) || (DLEQProofCKeys[i].a2.Y != pt2.Y))
            {
                VerificationResult.push(false);
                return false;
            }
        }
        VerificationResult.push(true);
        return true;
    }

    function DLEQVerifyDis(uint index) public payable returns (bool)
    {
        G1Point memory gG = g1mul(g, DLEQProofDis.z);
        G1Point memory y1G = g1mul(UserPk, DLEQProofDis.c);
        G1Point memory hG = g1mul(EKeys[index].EK0, DLEQProofDis.z);
        G1Point memory y2G = g1mul(Dispute, DLEQProofDis.c);

        G1Point memory pt1 =  g1add(gG, y1G);
        G1Point memory pt2 =  g1add(hG, y2G);
        if ((DLEQProofDis.a1.X != pt1.X) || (DLEQProofDis.a1.Y != pt1.Y) || (DLEQProofDis.a2.X != pt2.X) || (DLEQProofDis.a2.Y != pt2.Y))
        {
            VerificationResult.push(false);
            return false;
        }
        VerificationResult.push(true);
        return true;
    }

    // Encoding of field elements is: X[0] * z + X[1]


    struct Ciphertext {
        G1Point C0;
        G1Point C1;
    }

    struct DLEQProof {
        G1Point a1;
        G1Point a2;
        uint256 c;
        uint256 z;
    }

    struct EKey {
        G1Point EK0;
        G1Point EK1;
    }

    G1Point g;
    G1Point OwnerPk;
    G1Point UserPk;
    G1Point[] TTPsPk;
    Ciphertext ciphertext;
    G1Point[] Gs;
    G1Point[] CKeys;
    DLEQProof[] DLEQProofCKeys;
    DLEQProof[] DLEQProofKeys;
    G1Point Dispute;
    DLEQProof DLEQProofDis;
    EKey[] EKeys;
    bool[] VerificationResult;
*/
    struct SystemKey {
	    G1Point Tau1;
	    G2Point Tau2 ;   	
    }

    struct UKey {
	    G1Point pk;
	    G2Point vk;   	
    }

    SystemKey[] PK;
    UKey OwnerKey;
    UKey UserKey;
    G1Point[] DRsKey;
    
    function UploadSystemKey(G1Point[] memory tau1, G2Point[] memory tau2) public {
        SystemKey memory pk;
    	for (uint i = 0; i < tau1.length; i++) {
            pk.Tau1=tau1[i];
            pk.Tau2=tau2[i];
            PK.push(pk);
        }
    }

    function UploadOwnerKey(G1Point memory pk,G2Point memory vk) public {
        OwnerKey.pk = pk;
        OwnerKey.vk = vk;
    }

    function UploadUserKey(G1Point memory pk,G2Point memory vk) public {
        UserKey.pk = pk;
        UserKey.vk = vk;
    }

    function UploadDRsKey(G1Point[] memory pkArray) public {
        for (uint i = 0; i < pkArray.length; i++) {
            	DRsKey.push(pkArray[i]);
        	}
    }

/*   

    function UploadCiphertext(G1Point memory c0, G1Point memory c1) public {
        ciphertext.C0 = c0;
        ciphertext.C1 = c1;
    }


    function UploadGs(G1Point[] memory gs) public {
        for (uint i = 0; i < gs.length; i++) {
            Gs.push(gs[i]);
        }
    }

    function UploadCKeys(G1Point[] memory ckeys) public {
        //G1Point[] memory CKeys = new G1Point[](ckeys.length);
        for (uint i = 0; i < ckeys.length; i++) {
            CKeys.push(ckeys[i]);
        }
    }

    function UploadDLEQProofCKeys(uint256[] memory c, G1Point[] memory a1, G1Point[] memory a2, uint256[] memory z) public {
        DLEQProof memory DLEQProofCKey;
        for (uint i = 0; i < c.length; i++) {
            DLEQProofCKey.c=c[i];
            DLEQProofCKey.a1=a1[i];
            DLEQProofCKey.a2=a2[i];
            DLEQProofCKey.z=z[i];
            DLEQProofCKeys.push(DLEQProofCKey);
        }
    }

    function UploadDLEQProofKeys(uint256[] memory _c, G1Point[] memory _a1, G1Point[] memory _a2, uint256[] memory _z) public {
        DLEQProof memory DLEQProofKey;
        for (uint i = 0; i < _c.length; i++) {
            DLEQProofKey.a1 = _a1[i];
            DLEQProofKey.a2 = _a2[i];
            DLEQProofKey.c = _c[i];
            DLEQProofKey.z = _z[i];
            DLEQProofKeys.push(DLEQProofKey);
        }
    }
 
    function UploadDispute(G1Point memory Dis) public {
        Dispute=Dis;
    }

    function UploadDisputeProof(uint256 c, G1Point memory a1, G1Point memory a2, uint256 z) public {
        DLEQProofDis.a1 = a1;
        DLEQProofDis.a2 = a2;
        DLEQProofDis.c = c;
        DLEQProofDis.z = z;
    }

    function UploadEKeys(G1Point[] memory EKeys0, G1Point[] memory EKeys1) public {
       EKey memory eKey;
        for (uint i = 0; i < EKeys0.length; i++) {
            eKey.EK0 = EKeys0[i];
            eKey.EK1 = EKeys1[i];
            EKeys.push(eKey);
        }
    }

    function GetVrfResult() public view returns (bool[] memory) {
        return VerificationResult;
    }
*/
}
