import {
  Typography,
  TextField,
  Button,
  Box,
  Modal,
  Grid2,
  Alert,
  FormControl,
  InputLabel,
  Select,
  MenuItem,
  SelectChangeEvent,
} from "@mui/material";
import { ChangeEvent, useEffect, useState } from "react";
import { useDispatch, useSelector } from "react-redux";
import { useNavigate } from "react-router";
import { IAddress } from "../../addresses/interfaces/IAddress";
import { updateCarrier, updateShipping } from "../../../Store/checkout/checkoutSlice";
import { ArrowBackIosNew } from "@mui/icons-material";
import { ICarrier } from "../../carriers/interfaces/ICarrier";
import { CarrierService } from "../../carriers/services/carrierService";

export const ShippingPage = () => {
  const dispatch = useDispatch();
  const navigate = useNavigate();
  const shippingData = useSelector((state: any) => state.checkout.shipping);
  const [error, seterror] = useState(false);
  const [shipping, setshipping] = useState<IAddress>(shippingData);
  const [carriers, setcarriers] = useState<ICarrier[]>([]);
  const [selectedCarrier, setselectedCarrier] = useState<ICarrier|null>(null)
  
  const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;

    setshipping({
      ...shipping,
      [name]: name === "number" ? Number(value) || 0 : value,
    });
  };

  const handleCarrierChange = (e: SelectChangeEvent<string>) => {
    const carrierID = Number(e.target.value);
    const selected = carriers.find((carrier) => carrier.ID === carrierID) || null;
    setselectedCarrier(selected);
  };

  //comprobacion
  useEffect(() => {
    console.log("VALOR DE SELECTED CARRIER",selectedCarrier)

  }, [selectedCarrier])
  
  const checkSelectedCarrier = (carrier: ICarrier): boolean => {
    if (!carrier.ID) {
      return false
    }
    if (
      carrier.contact.trim() === "" ||
      carrier.name.trim() === "" ||
      carrier.ID <= 0
    ) {
      return false;
    }
    return true;
  };
  const checkShipping = (address: IAddress): boolean => {
    if (
      address.city.trim() === "" ||
      address.country.trim() === "" ||
      address.number < 0 ||
      address.postalCode.trim() === "" ||
      address.province.trim() === "" ||
      address.street.trim() === ""
    ) {
      return false;
    }
    return true;
  };
  const handleNext = () => {
    if (!selectedCarrier||!shipping) {
      console.log("No se han rellenado los datos de shipping")
      return
    }
    const ship: IAddress = { ...shipping };
    const carr: ICarrier = { ...selectedCarrier };
    if (checkShipping(ship)&&checkSelectedCarrier(carr)) {
      dispatch(updateShipping(ship));
      dispatch(updateCarrier(carr));
      navigate("/checkout/payment");
      return;
    }
    seterror(true);
    setTimeout(() => {
      seterror(false);
    }, 3000);
  };
  const handleGoBack = () => {
    navigate(-1);
  };

  const fetchCarriers = async () => {
    const response = await CarrierService.GetCarriers();
    if (response.success && response.data) {
      setcarriers(response.data);
      console.log(response.data);
    } else {
      seterror(true)
    }
  };

  useEffect(() => {
    fetchCarriers();
  }, []);

  const modalStyle = {
    position: "absolute",
    top: "50%",
    left: "50%",
    transform: "translate(-50%, -50%)",
    width: 400,
    bgcolor: "background.paper",
    border: "2px solid #000",
    boxShadow: 24,
    p: 4,
  };
  return (
    <div>
      {/* Modal de Shipping */}
      <Modal
        open={true} // Se abrirá por defecto, esto es solo para fines de demostración
        onClose={() => {}}
        aria-labelledby="shipping-modal-title"
        aria-describedby="shipping-modal-description"
      >
        <Box sx={modalStyle}>
          <Button
            onClick={handleGoBack}
            startIcon={<ArrowBackIosNew />}
            sx={{ marginBottom: 3 }}
          >
            Volver
          </Button>
          <Typography id="shipping-modal-title" variant="h6" component="h2">
            Shipping Information
          </Typography>
          <Typography id="shipping-modal-description" sx={{ mt: 2, mb: 3 }}>
            Please enter your shipping details below:
          </Typography>

          <Grid2
            container
            spacing={2}
            rowSpacing={6}
            columnSpacing={6}
            direction="row"
            sx={{ justifyContent: "space-between", alignItems: "center" }}
          >
            <Grid2 size={{ xs: 12, sm: 6 }}>
              <TextField
                onChange={handleChange}
                value={shipping.country}
                name="country"
                fullWidth
                label="Country"
                variant="outlined"
              />
            </Grid2>
            <Grid2 size={{ xs: 12, sm: 6 }}>
              <TextField
                onChange={handleChange}
                name="province"
                value={shipping.province}
                fullWidth
                label="Province"
                variant="outlined"
              />
            </Grid2>
            <Grid2 size={{ xs: 12, sm: 6 }}>
              <TextField
                onChange={handleChange}
                name="city"
                value={shipping.city}
                fullWidth
                label="City"
                variant="outlined"
              />
            </Grid2>
            <Grid2 size={{ xs: 12, sm: 6 }}>
              <TextField
                onChange={handleChange}
                name="postalCode"
                value={shipping.postalCode}
                fullWidth
                label="Postal code"
                variant="outlined"
              />
            </Grid2>
            <Grid2 size={{ xs: 12, sm: 6 }}>
              <TextField
                onChange={handleChange}
                name="street"
                value={shipping.street}
                fullWidth
                label="Street"
                variant="outlined"
              />
            </Grid2>
            <Grid2 size={{ xs: 12, sm: 6 }}>
              <TextField
                onChange={handleChange}
                name="number"
                value={shipping.number}
                fullWidth
                label="Number"
                variant="outlined"
                type="number"
              />
            </Grid2>

            <FormControl fullWidth>
              <InputLabel>Selecciona un transportista</InputLabel>
              <Select onChange={handleCarrierChange} defaultValue="">
                <MenuItem disabled value="">
                  ---
                </MenuItem>
                {carriers.map((carrier) => (
                  <MenuItem key={carrier.ID} value={carrier.ID}>
                    {carrier.name}
                  </MenuItem>
                ))}
              </Select>
            </FormControl>
          </Grid2>

          <Box sx={{ mt: 3 }}>
            <Button
              onClick={handleNext}
              variant="contained"
              color="primary"
              fullWidth
            >
              Confirm Shipping
            </Button>
          </Box>
          <Box sx={{ mt: 1 }}>
            {error && <Alert severity="error">Faltan datos por rellenar</Alert>}
          </Box>
        </Box>
      </Modal>
    </div>
  );
};
