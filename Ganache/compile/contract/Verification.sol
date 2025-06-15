// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

// https://www.iacr.org/cryptodb/archive/2002/ASIACRYPT/50/50.pdf
contract Verification
{
    uint256 constant FIELD_ORDER = 0x30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47;
    uint256 constant CURVE_ORDER = 0x30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001;
    uint256 constant GROUP_ORDER   = 21888242871839275222246405745257275088548364400416034343698204186575808495617;
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

	/// return the generator of G2
	function P2() pure internal returns (G2Point memory) {
		return G2Point(
			[11559732032986387107991004021392285783925812861821192530917403151452391805634,
			 10857046999023057135944570762232829481370756359578518086990519993285655852781],
			[4082367875863433681332203403145435568316851327593401208105741076214120093531,
			 8495653923123431417604973247489272438418190587263600148770280649306958101930]
		);
	}
    /// return the negation of p, i.e. p.add(p.negate()) should be zero.
	function g1neg(G1Point memory p) pure internal returns (G1Point memory) {
		// The prime q in the base field F_q for G1
		uint q = 21888242871839275222246405745257275088696311157297823662689037894645226208583;
		if (p.X == 0 && p.Y == 0)
			return G1Point(0, 0);
		return G1Point(p.X, q - (p.Y % q));
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


    
    //Operation in group G2
    uint256 internal constant FIELD_MODULUS = 0x30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47;
    uint256 internal constant TWISTBX = 0x2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5;
    uint256 internal constant TWISTBY = 0x9713b03af0fed4cd2cafadeed8fdf4a74fa084e52d1852e4a2bd0685c315d2;
    //266929791119991161246907387137283842545076965332900288569378510910307636690
    uint internal constant PTXX = 0;
    uint internal constant PTXY = 1;
    uint internal constant PTYX = 2;
    uint internal constant PTYY = 3;
    uint internal constant PTZX = 4;
    uint internal constant PTZY = 5;

    /**
     * @notice Add two twist points
     * @param pt1xx Coefficient 1 of x on point 1
     * @param pt1xy Coefficient 2 of x on point 1
     * @param pt1yx Coefficient 1 of y on point 1
     * @param pt1yy Coefficient 2 of y on point 1
     * @param pt2xx Coefficient 1 of x on point 2
     * @param pt2xy Coefficient 2 of x on point 2
     * @param pt2yx Coefficient 1 of y on point 2
     * @param pt2yy Coefficient 2 of y on point 2
     * @return (pt3xx, pt3xy, pt3yx, pt3yy)
     */
    function ECTwistAdd(
        uint256 pt1xx, uint256 pt1xy,
        uint256 pt1yx, uint256 pt1yy,
        uint256 pt2xx, uint256 pt2xy,
        uint256 pt2yx, uint256 pt2yy
    ) view internal returns (
        uint256, uint256,
        uint256, uint256
    ) {
        if (
            pt1xx == 0 && pt1xy == 0 &&
            pt1yx == 0 && pt1yy == 0
        ) {
            if (!(
                pt2xx == 0 && pt2xy == 0 &&
                pt2yx == 0 && pt2yy == 0
            )) {
                assert(_isOnCurve(
                    pt2xx, pt2xy,
                    pt2yx, pt2yy
                ));
            }
            return (
                pt2xx, pt2xy,
                pt2yx, pt2yy
            );
        } else if (
            pt2xx == 0 && pt2xy == 0 &&
            pt2yx == 0 && pt2yy == 0
        ) {
            assert(_isOnCurve(
                pt1xx, pt1xy,
                pt1yx, pt1yy
            ));
            return (
                pt1xx, pt1xy,
                pt1yx, pt1yy
            );
        }

        assert(_isOnCurve(
            pt1xx, pt1xy,
            pt1yx, pt1yy
        ));
        assert(_isOnCurve(
            pt2xx, pt2xy,
            pt2yx, pt2yy
        ));

        uint256[6] memory pt3 = _ECTwistAddJacobian(
            pt1xx, pt1xy,
            pt1yx, pt1yy,
            1,     0,
            pt2xx, pt2xy,
            pt2yx, pt2yy,
            1,     0
        );

        return _fromJacobian(
            pt3[PTXX], pt3[PTXY],
            pt3[PTYX], pt3[PTYY],
            pt3[PTZX], pt3[PTZY]
        );
    }

    /**
     * @notice Multiply a twist point by a scalar
     * @param s     Scalar to multiply by
     * @param pt1xx Coefficient 1 of x
     * @param pt1xy Coefficient 2 of x
     * @param pt1yx Coefficient 1 of y
     * @param pt1yy Coefficient 2 of y
     * @return (pt2xx, pt2xy, pt2yx, pt2yy)
     */
    function ECTwistMul(
        uint256 s,
        uint256 pt1xx, uint256 pt1xy,
        uint256 pt1yx, uint256 pt1yy
    ) view internal returns (
        uint256, uint256,
        uint256, uint256
    ) {
        uint256 pt1zx = 1;
        if (
            pt1xx == 0 && pt1xy == 0 &&
            pt1yx == 0 && pt1yy == 0
        ) {
            pt1xx = 1;
            pt1yx = 1;
            pt1zx = 0;
        } else {
            assert(_isOnCurve(
                pt1xx, pt1xy,
                pt1yx, pt1yy
            ));
        }

        uint256[6] memory pt2 = _ECTwistMulJacobian(
            s,
            pt1xx, pt1xy,
            pt1yx, pt1yy,
            pt1zx, 0
        );

        return _fromJacobian(
            pt2[PTXX], pt2[PTXY],
            pt2[PTYX], pt2[PTYY],
            pt2[PTZX], pt2[PTZY]
        );
    }

    /**
     * @notice Get the field modulus
     * @return The field modulus
     */

    function ECTwistNeg(uint256 pt1xx, uint256 pt1xy, uint256 pt1yx, uint256 pt1yy) pure internal returns (uint256, uint256,uint256, uint256) {
		// The prime q in the base field F_q for G2
		uint q = 21888242871839275222246405745257275088696311157297823662689037894645226208583;
		if (pt1xx == 0 && pt1xy == 0 && pt1yx == 0 && pt1yy == 0)
			return (0, 0, 0, 0);
		return (pt1xx,pt1xy, q - (pt1yx % q), q - (pt1yy % q));
	}



    function GetFieldModulus() public pure returns (uint256) {
        return FIELD_MODULUS;
    }

    function submod(uint256 a, uint256 b, uint256 n) internal pure returns (uint256) {
        return addmod(a, n - b, n);
    }

    function _FQ2Mul(
        uint256 xx, uint256 xy,
        uint256 yx, uint256 yy
    ) internal pure returns (uint256, uint256) {
        return (
            submod(mulmod(xx, yx, FIELD_MODULUS), mulmod(xy, yy, FIELD_MODULUS), FIELD_MODULUS),
            addmod(mulmod(xx, yy, FIELD_MODULUS), mulmod(xy, yx, FIELD_MODULUS), FIELD_MODULUS)
        );
    }

    function _FQ2Muc(
        uint256 xx, uint256 xy,
        uint256 c
    ) internal pure returns (uint256, uint256) {
        return (
            mulmod(xx, c, FIELD_MODULUS),
            mulmod(xy, c, FIELD_MODULUS)
        );
    }

    function _FQ2Add(
        uint256 xx, uint256 xy,
        uint256 yx, uint256 yy
    ) internal pure returns (uint256, uint256) {
        return (
            addmod(xx, yx, FIELD_MODULUS),
            addmod(xy, yy, FIELD_MODULUS)
        );
    }

    function _FQ2Sub(
        uint256 xx, uint256 xy,
        uint256 yx, uint256 yy
    ) internal pure returns (uint256 rx, uint256 ry) {
        return (
            submod(xx, yx, FIELD_MODULUS),
            submod(xy, yy, FIELD_MODULUS)
        );
    }

    function _FQ2Div(
        uint256 xx, uint256 xy,
        uint256 yx, uint256 yy
    ) internal view returns (uint256, uint256) {
        (yx, yy) = _FQ2Inv(yx, yy);
        return _FQ2Mul(xx, xy, yx, yy);
    }

    function _FQ2Inv(uint256 x, uint256 y) internal view returns (uint256, uint256) {
        uint256 inv = _modInv(addmod(mulmod(y, y, FIELD_MODULUS), mulmod(x, x, FIELD_MODULUS), FIELD_MODULUS), FIELD_MODULUS);
        return (
            mulmod(x, inv, FIELD_MODULUS),
            FIELD_MODULUS - mulmod(y, inv, FIELD_MODULUS)
        );
    }

    function _isOnCurve(
        uint256 xx, uint256 xy,
        uint256 yx, uint256 yy
    ) internal pure returns (bool) {
        uint256 yyx;
        uint256 yyy;
        uint256 xxxx;
        uint256 xxxy;
        (yyx, yyy) = _FQ2Mul(yx, yy, yx, yy);
        (xxxx, xxxy) = _FQ2Mul(xx, xy, xx, xy);
        (xxxx, xxxy) = _FQ2Mul(xxxx, xxxy, xx, xy);
        (yyx, yyy) = _FQ2Sub(yyx, yyy, xxxx, xxxy);
        (yyx, yyy) = _FQ2Sub(yyx, yyy, TWISTBX, TWISTBY);
        return yyx == 0 && yyy == 0;
    }

    function _modInv(uint256 a, uint256 n) internal view returns (uint256 result) {
        bool success;
        assembly {
            let freemem := mload(0x40)
            mstore(freemem, 0x20)
            mstore(add(freemem,0x20), 0x20)
            mstore(add(freemem,0x40), 0x20)
            mstore(add(freemem,0x60), a)
            mstore(add(freemem,0x80), sub(n, 2))
            mstore(add(freemem,0xA0), n)
            success := staticcall(sub(gas(), 2000), 5, freemem, 0xC0, freemem, 0x20)
            result := mload(freemem)
        }
        require(success);
    }

    function _fromJacobian(
        uint256 pt1xx, uint256 pt1xy,
        uint256 pt1yx, uint256 pt1yy,
        uint256 pt1zx, uint256 pt1zy
    ) internal view returns (
        uint256 pt2xx, uint256 pt2xy,
        uint256 pt2yx, uint256 pt2yy
    ) {
        uint256 invzx;
        uint256 invzy;
        (invzx, invzy) = _FQ2Inv(pt1zx, pt1zy);
        (pt2xx, pt2xy) = _FQ2Mul(pt1xx, pt1xy, invzx, invzy);
        (pt2yx, pt2yy) = _FQ2Mul(pt1yx, pt1yy, invzx, invzy);
    }

    function _ECTwistAddJacobian(
        uint256 pt1xx, uint256 pt1xy,
        uint256 pt1yx, uint256 pt1yy,
        uint256 pt1zx, uint256 pt1zy,
        uint256 pt2xx, uint256 pt2xy,
        uint256 pt2yx, uint256 pt2yy,
        uint256 pt2zx, uint256 pt2zy) internal pure returns (uint256[6] memory pt3) {
            if (pt1zx == 0 && pt1zy == 0) {
                (
                    pt3[PTXX], pt3[PTXY],
                    pt3[PTYX], pt3[PTYY],
                    pt3[PTZX], pt3[PTZY]
                ) = (
                    pt2xx, pt2xy,
                    pt2yx, pt2yy,
                    pt2zx, pt2zy
                );
                return pt3;
            } else if (pt2zx == 0 && pt2zy == 0) {
                (
                    pt3[PTXX], pt3[PTXY],
                    pt3[PTYX], pt3[PTYY],
                    pt3[PTZX], pt3[PTZY]
                ) = (
                    pt1xx, pt1xy,
                    pt1yx, pt1yy,
                    pt1zx, pt1zy
                );
                return pt3;
            }

            (pt2yx,     pt2yy)     = _FQ2Mul(pt2yx, pt2yy, pt1zx, pt1zy); // U1 = y2 * z1
            (pt3[PTYX], pt3[PTYY]) = _FQ2Mul(pt1yx, pt1yy, pt2zx, pt2zy); // U2 = y1 * z2
            (pt2xx,     pt2xy)     = _FQ2Mul(pt2xx, pt2xy, pt1zx, pt1zy); // V1 = x2 * z1
            (pt3[PTZX], pt3[PTZY]) = _FQ2Mul(pt1xx, pt1xy, pt2zx, pt2zy); // V2 = x1 * z2

            if (pt2xx == pt3[PTZX] && pt2xy == pt3[PTZY]) {
                if (pt2yx == pt3[PTYX] && pt2yy == pt3[PTYY]) {
                    (
                        pt3[PTXX], pt3[PTXY],
                        pt3[PTYX], pt3[PTYY],
                        pt3[PTZX], pt3[PTZY]
                    ) = _ECTwistDoubleJacobian(pt1xx, pt1xy, pt1yx, pt1yy, pt1zx, pt1zy);
                    return pt3;
                }
                (
                    pt3[PTXX], pt3[PTXY],
                    pt3[PTYX], pt3[PTYY],
                    pt3[PTZX], pt3[PTZY]
                ) = (
                    1, 0,
                    1, 0,
                    0, 0
                );
                return pt3;
            }

            (pt2zx,     pt2zy)     = _FQ2Mul(pt1zx, pt1zy, pt2zx,     pt2zy);     // W = z1 * z2
            (pt1xx,     pt1xy)     = _FQ2Sub(pt2yx, pt2yy, pt3[PTYX], pt3[PTYY]); // U = U1 - U2
            (pt1yx,     pt1yy)     = _FQ2Sub(pt2xx, pt2xy, pt3[PTZX], pt3[PTZY]); // V = V1 - V2
            (pt1zx,     pt1zy)     = _FQ2Mul(pt1yx, pt1yy, pt1yx,     pt1yy);     // V_squared = V * V
            (pt2yx,     pt2yy)     = _FQ2Mul(pt1zx, pt1zy, pt3[PTZX], pt3[PTZY]); // V_squared_times_V2 = V_squared * V2
            (pt1zx,     pt1zy)     = _FQ2Mul(pt1zx, pt1zy, pt1yx,     pt1yy);     // V_cubed = V * V_squared
            (pt3[PTZX], pt3[PTZY]) = _FQ2Mul(pt1zx, pt1zy, pt2zx,     pt2zy);     // newz = V_cubed * W
            (pt2xx,     pt2xy)     = _FQ2Mul(pt1xx, pt1xy, pt1xx,     pt1xy);     // U * U
            (pt2xx,     pt2xy)     = _FQ2Mul(pt2xx, pt2xy, pt2zx,     pt2zy);     // U * U * W
            (pt2xx,     pt2xy)     = _FQ2Sub(pt2xx, pt2xy, pt1zx,     pt1zy);     // U * U * W - V_cubed
            (pt2zx,     pt2zy)     = _FQ2Muc(pt2yx, pt2yy, 2);                    // 2 * V_squared_times_V2
            (pt2xx,     pt2xy)     = _FQ2Sub(pt2xx, pt2xy, pt2zx,     pt2zy);     // A = U * U * W - V_cubed - 2 * V_squared_times_V2
            (pt3[PTXX], pt3[PTXY]) = _FQ2Mul(pt1yx, pt1yy, pt2xx,     pt2xy);     // newx = V * A
            (pt1yx,     pt1yy)     = _FQ2Sub(pt2yx, pt2yy, pt2xx,     pt2xy);     // V_squared_times_V2 - A
            (pt1yx,     pt1yy)     = _FQ2Mul(pt1xx, pt1xy, pt1yx,     pt1yy);     // U * (V_squared_times_V2 - A)
            (pt1xx,     pt1xy)     = _FQ2Mul(pt1zx, pt1zy, pt3[PTYX], pt3[PTYY]); // V_cubed * U2
            (pt3[PTYX], pt3[PTYY]) = _FQ2Sub(pt1yx, pt1yy, pt1xx,     pt1xy);     // newy = U * (V_squared_times_V2 - A) - V_cubed * U2
    }

    function _ECTwistDoubleJacobian(
        uint256 pt1xx, uint256 pt1xy,
        uint256 pt1yx, uint256 pt1yy,
        uint256 pt1zx, uint256 pt1zy
    ) internal pure returns (
        uint256 pt2xx, uint256 pt2xy,
        uint256 pt2yx, uint256 pt2yy,
        uint256 pt2zx, uint256 pt2zy
    ) {
        (pt2xx, pt2xy) = _FQ2Muc(pt1xx, pt1xy, 3);            // 3 * x
        (pt2xx, pt2xy) = _FQ2Mul(pt2xx, pt2xy, pt1xx, pt1xy); // W = 3 * x * x
        (pt1zx, pt1zy) = _FQ2Mul(pt1yx, pt1yy, pt1zx, pt1zy); // S = y * z
        (pt2yx, pt2yy) = _FQ2Mul(pt1xx, pt1xy, pt1yx, pt1yy); // x * y
        (pt2yx, pt2yy) = _FQ2Mul(pt2yx, pt2yy, pt1zx, pt1zy); // B = x * y * S
        (pt1xx, pt1xy) = _FQ2Mul(pt2xx, pt2xy, pt2xx, pt2xy); // W * W
        (pt2zx, pt2zy) = _FQ2Muc(pt2yx, pt2yy, 8);            // 8 * B
        (pt1xx, pt1xy) = _FQ2Sub(pt1xx, pt1xy, pt2zx, pt2zy); // H = W * W - 8 * B
        (pt2zx, pt2zy) = _FQ2Mul(pt1zx, pt1zy, pt1zx, pt1zy); // S_squared = S * S
        (pt2yx, pt2yy) = _FQ2Muc(pt2yx, pt2yy, 4);            // 4 * B
        (pt2yx, pt2yy) = _FQ2Sub(pt2yx, pt2yy, pt1xx, pt1xy); // 4 * B - H
        (pt2yx, pt2yy) = _FQ2Mul(pt2yx, pt2yy, pt2xx, pt2xy); // W * (4 * B - H)
        (pt2xx, pt2xy) = _FQ2Muc(pt1yx, pt1yy, 8);            // 8 * y
        (pt2xx, pt2xy) = _FQ2Mul(pt2xx, pt2xy, pt1yx, pt1yy); // 8 * y * y
        (pt2xx, pt2xy) = _FQ2Mul(pt2xx, pt2xy, pt2zx, pt2zy); // 8 * y * y * S_squared
        (pt2yx, pt2yy) = _FQ2Sub(pt2yx, pt2yy, pt2xx, pt2xy); // newy = W * (4 * B - H) - 8 * y * y * S_squared
        (pt2xx, pt2xy) = _FQ2Muc(pt1xx, pt1xy, 2);            // 2 * H
        (pt2xx, pt2xy) = _FQ2Mul(pt2xx, pt2xy, pt1zx, pt1zy); // newx = 2 * H * S
        (pt2zx, pt2zy) = _FQ2Mul(pt1zx, pt1zy, pt2zx, pt2zy); // S * S_squared
        (pt2zx, pt2zy) = _FQ2Muc(pt2zx, pt2zy, 8);            // newz = 8 * S * S_squared
    }

    function _ECTwistMulJacobian(
        uint256 d,
        uint256 pt1xx, uint256 pt1xy,
        uint256 pt1yx, uint256 pt1yy,
        uint256 pt1zx, uint256 pt1zy
    ) internal pure returns (uint256[6] memory pt2) {
        while (d != 0) {
            if ((d & 1) != 0) {
                pt2 = _ECTwistAddJacobian(
                    pt2[PTXX], pt2[PTXY],
                    pt2[PTYX], pt2[PTYY],
                    pt2[PTZX], pt2[PTZY],
                    pt1xx, pt1xy,
                    pt1yx, pt1yy,
                    pt1zx, pt1zy);
            }
            (
                pt1xx, pt1xy,
                pt1yx, pt1yy,
                pt1zx, pt1zy
            ) = _ECTwistDoubleJacobian(
                pt1xx, pt1xy,
                pt1yx, pt1yy,
                pt1zx, pt1zy
            );

            d = d / 2;
        }
    }

	function pairing(G1Point[] memory p1, G2Point[] memory p2) view internal returns (bool) {
		require(p1.length == p2.length);
		uint elements = p1.length;
		uint inputSize = elements * 6;
		uint[] memory input = new uint[](inputSize);
		for (uint i = 0; i < elements; i++)
		{
			input[i * 6 + 0] = p1[i].X;
			input[i * 6 + 1] = p1[i].Y;
			input[i * 6 + 2] = p2[i].X[0];
			input[i * 6 + 3] = p2[i].X[1];
			input[i * 6 + 4] = p2[i].Y[0];
			input[i * 6 + 5] = p2[i].Y[1];
		}
		uint[1] memory out;
		bool success;
		assembly {
			success := staticcall(sub(gas()	, 2000), 8, add(input, 0x20), mul(inputSize, 0x20), out, 0x20)
			// Use "invalid" to make gas estimation work
			//switch success case 0 { invalid }
		}
		require(success);
		return out[0] != 0;
	}


	/// Convenience method for a pairing check for two pairs.
	function pairingProd2(G1Point memory a1, G2Point memory a2, G1Point memory b1, G2Point memory b2) view internal returns (bool) {
		G1Point[] memory p1 = new G1Point[](2);
		G2Point[] memory p2 = new G2Point[](2);
		p1[0] = a1;
		p1[1] = b1;
		p2[0] = a2;
		p2[1] = b2;
		return pairing(p1, p2);
	}

	/// Convenience method for a pairing check for three pairs.
	function pairingProd3(
			G1Point memory a1, G2Point memory a2,
			G1Point memory b1, G2Point memory b2,
			G1Point memory c1, G2Point memory c2
	) view internal returns (bool) {
		G1Point[] memory p1 = new G1Point[](3);
		G2Point[] memory p2 = new G2Point[](3);
		p1[0] = a1;
		p1[1] = b1;
		p1[2] = c1;
		p2[0] = a2;
		p2[1] = b2;
		p2[2] = c2;
		return pairing(p1, p2);
	}

	/// Convenience method for a pairing check for four pairs.
	function pairingProd4(
			G1Point memory a1, G2Point memory a2,
			G1Point memory b1, G2Point memory b2,
			G1Point memory c1, G2Point memory c2,
			G1Point memory d1, G2Point memory d2
	) view internal returns (bool) {
		G1Point[] memory p1 = new G1Point[](4);
		G2Point[] memory p2 = new G2Point[](4);
		p1[0] = a1;
		p1[1] = b1;
		p1[2] = c1;
		p1[3] = d1;
		p2[0] = a2;
		p2[1] = b2;
		p2[2] = c2;
		p2[3] = d2;
		return pairing(p1, p2);
	}

    function DLEQVerify(G1Point memory g, G1Point memory y1, G1Point memory a1, G1Point memory h, G1Point memory y2, G1Point memory a2, uint256 c, uint256 z) public payable returns (bool)
    {
        G1Point memory gG = g1mul(g, z);
        G1Point memory y1G = g1mul(y1, c);

        G1Point memory hG = g1mul(h, z);
        G1Point memory y2G = g1mul(y2, c);

        G1Point memory pt1 =  g1add(gG, y1G);
        G1Point memory pt2 =  g1add(hG, y2G);
        if ((a1.X != pt1.X) || (a1.Y != pt1.Y) || (a2.X != pt2.X) || (a2.Y != pt2.Y))
        {
            return false;
        }
        return true;
    }

/*
    mapping(uint256 => uint256) public invMap;
    bool RScodeResult;

    //RScode verify on-chain
    function RScodeVerify() public
    {
        G1Point memory sum=PK[0].Tau1;
        G1Point memory codeword;
        uint len = ReKeys.length;
        uint256 result=1;
        for(uint i=1;i< ReKeys.length;i++)
        {
            result = 1;
            for(uint j=1; j< ReKeys.length;j++)
            {
                if(i!=j)
                {
                    result=mulmod(result, invMap[i+GROUP_ORDER-j], GROUP_ORDER);
                }
            }
            codeword = g1mul(ReKeys[i-1].RK0,result);
            sum=g1add(sum,codeword);
        }
        if(sum.X==PK[0].Tau1.X && sum.Y== PK[0].Tau1.Y)
        {
            RScodeResult=true;
        }
        else
        {
            RScodeResult=false;
        }
    }


    function GetRScodeResult() public view returns (bool){
        return RScodeResult;
    }
*/


    struct SystemKey {
	    G1Point Tau1;
	    G2Point Tau2;   	
    }

    struct UKey {
	    G1Point pk;
	    G2Point vk;   	
    }

    struct G1DLEQProof {
        G1Point a1;
        G1Point a2;
        uint256 c;
        uint256 z;
    }

    struct G1G2DLEQProof {
        G1Point a1;
        G2Point a2;
        uint256 c;
        uint256 z;
    }

    struct Cipher {
        G2Point C0;
        string C1;
    }

    struct ReKey{
        G1Point RK0;
        G1Point RK1;
    }

    struct ReKeyProof{
        G1Point Witness;
        G1DLEQProof ShareDLEQ;
    }

    struct ReCipher{
        G1Point C2;
        G1Point C3;
    }


    SystemKey[] PK;
    UKey OwnerKey;
    UKey UserKey;
    G1Point[] DRsKey;
    G1G2DLEQProof VKoDLEQ;
    G1G2DLEQProof VKuDLEQ;
    bool VKoResult;
    bool VKuResult;
    Cipher ciphertext;
    ReKey[] ReKeys;
    G1Point Commitment;
    ReKeyProof[] ReKeysProof;
    bool[] ReKeysResult;
    ReCipher[] reciphertexts;
    uint256[] index;

    function Empty(G2Point memory a, G2Point memory b)public{
        return;
    }

    function Add(G2Point memory a, G2Point memory b)public{
        G2Point memory sum;
        (sum.X[1],sum.X[0],sum.Y[1],sum.Y[0])=ECTwistAdd(a.X[1],a.X[0],a.Y[1],a.Y[0],b.X[1],b.X[0],b.Y[1],b.Y[0]);
        return;
    }

    function Mul(G2Point memory a, uint256 b)public{
       G2Point memory sum;
       (sum.X[1],sum.X[0],sum.Y[1],sum.Y[0])=ECTwistMul(b,a.X[1],a.X[0],a.Y[1],a.Y[0]);
       return;
    }

    
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


    function VKoVerify() public 
    {
        if (pairingProd2(PK[0].Tau1, OwnerKey.vk, g1neg(OwnerKey.pk), PK[0].Tau2))
        {
            VKoResult=true;
        }
        else
        {
            VKoResult=false;
        }    
    }
    function GetVKoResult() public view returns (bool) {
        return VKoResult;
    }

    function VKuVerify() public
    {
        if (pairingProd2(PK[0].Tau1, UserKey.vk, g1neg(UserKey.pk), PK[0].Tau2))
        {
            VKuResult=true;
        }
        else
        {
            VKuResult=false;
        }      
    }
    function GetVKuResult() public view returns (bool) {
        return VKuResult;
    }


    function UploadCiphertext(G2Point memory c0, string memory c1) public {
        ciphertext.C0 = c0;
        ciphertext.C1 = c1;
    }

    function UploadReKeys(G1Point[] memory rks0, G1Point[] memory rks1) public {
        ReKey memory rk;
        for (uint i = 0; i < rks0.length; i++) {
            rk.RK0 = rks0[i];
            rk.RK1 = rks1[i];
            ReKeys.push(rk);
        }
    }

    function UploadReKeysProof(uint256[] memory I, G1Point memory commitment, G1Point[] memory witness, G1Point[] memory a1, G1Point[] memory a2,uint256[] memory c, uint256[] memory z)public{
        for (uint i=0;i<I.length;i++){
            index.push(I[i]);
        }
        Commitment=commitment;
        G1DLEQProof memory prf;
        ReKeyProof memory proof;
        for (uint i = 0; i < witness.length; i++) {
            prf.a1=a1[i];
            prf.a2=a2[i];
            prf.c=c[i];
            prf.z=z[i];
            proof.Witness = witness[i];
            proof.ShareDLEQ = prf;
            ReKeysProof.push(proof);
        }
    }

   function ReKeysVerify()public{
        G1Point memory base;
        G1Point memory left;
        G2Point memory g2i;
        G2Point memory right;

        for (uint256 i = 0; i < ReKeysProof.length; i++) 
        {
            //Verify e(commitment/g^{s_i},g2)e((witness[i])^{-1},g2^{\tau}-g2^i)==1
            left = g1add(Commitment,g1neg(ReKeys[i].RK0));
            (g2i.X[1],g2i.X[0],g2i.Y[1],g2i.Y[0]) = ECTwistMul(uint256(index[i]),PK[0].Tau2.X[1],PK[0].Tau2.X[0],PK[0].Tau2.Y[1],PK[0].Tau2.Y[0]);
            (g2i.X[1],g2i.X[0],g2i.Y[1],g2i.Y[0])=ECTwistNeg(g2i.X[1],g2i.X[0],g2i.Y[1],g2i.Y[0]);
            (right.X[1],right.X[0],right.Y[1],right.Y[0]) = ECTwistAdd(PK[1].Tau2.X[1],PK[1].Tau2.X[0],PK[1].Tau2.Y[1],PK[1].Tau2.Y[0],g2i.X[1],g2i.X[0],g2i.Y[1],g2i.Y[0]);
            if (pairingProd2(left, PK[0].Tau2, g1neg(ReKeysProof[i].Witness), right))
            {
                base = g1add(OwnerKey.pk,g1add(DRsKey[i],UserKey.pk));
                if (!DLEQVerify(PK[0].Tau1,ReKeys[i].RK0,ReKeysProof[i].ShareDLEQ.a1,base,ReKeys[i].RK1,ReKeysProof[i].ShareDLEQ.a2,ReKeysProof[i].ShareDLEQ.c,ReKeysProof[i].ShareDLEQ.z)){
                    index[i]=0;
                }
                ReKeysResult.push(true);
            }
            else
            {
                index[i]=0;
            }
        }
    }

    function GetReKeysResult() public view returns (bool[] memory){
        return ReKeysResult;
    }


    function UploadReCipher(G1Point[] memory c3) public{
        ReCipher memory recipher;
        for (uint i=0;i<index.length;i++){
            recipher.C2= ReKeys[index[i]-1].RK0;
            recipher.C3=c3[i];
            reciphertexts.push(recipher);
        }
    }

    function ReCipherVerify()public{
        for (uint i=0;i<reciphertexts.length;i++){
            if (!pairingProd3(g1neg(reciphertexts[i].C3),PK[0].Tau2,reciphertexts[i].C2,OwnerKey.vk,reciphertexts[i].C2,UserKey.vk)){
                index[i]=0;
            }
        }
    }


    function GetIndex() public view returns (uint[] memory){
        return index;
    }

    function GetReCipher() public view returns (G2Point memory, string memory,ReCipher[] memory){
        return (ciphertext.C0,ciphertext.C1,reciphertexts);
    }

}
